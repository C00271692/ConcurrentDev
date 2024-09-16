#include <iostream>
int romanToArabic(std::string romanNum)
{
    int arabicNum = 0;
    int I = 1;
    int V = 5;
    int X = 10;
    //int L = 50;
    //int C = 100;
    //int D = 500;
    //int M = 1000;

    for(int i = 0; i < romanNum.length(); i++)
    {
        if(romanNum[i] == 'I')
        {
            if(i + 1 < romanNum.length() && (romanNum[i + 1] == 'V' || romanNum[i + 1] == 'X'))
            {
                arabicNum-=I;
            }
            else
            {
                arabicNum+=I;
            }
        }
        if(romanNum[i] == 'V')
        {
            arabicNum+=V;
        }
        if(romanNum[i] == 'X')
        {
            arabicNum+=X;
        }

    }




    return arabicNum;
}


int main()
{
    std::string romanNumInput;
    std::cout << "Input Roman numeral: ";
    std::cin >> romanNumInput;
    int arabicNum = romanToArabic(romanNumInput); // Convert to Arabic numeral
    std::cout << "Arabic numeral: " << arabicNum << std::endl; // Output the result
    return 0;
}
