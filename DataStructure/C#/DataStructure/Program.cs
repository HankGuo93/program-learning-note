using System;
using System.Linq;

namespace DataStructure
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var heap = new MyHeap();

            for (int i = 0; i < 10; i++)
            {
                var random = new Random();
                var randomNum = random.Next(1, 10);
                heap.Insert(randomNum);
            }

            Console.WriteLine("print arraydata : \n");
            for (int i = 0; i < heap.Counter; i++)
            {
                Console.WriteLine(heap.ArrayData[i]);
            }

            Console.WriteLine("\n---------------------\n");

            var counter = heap.Counter;
            Console.WriteLine("print arraydata (from top) : \n");
            for (int i = 0; i < counter; i++)
            {
                Console.WriteLine(string.Join("-", heap.ArrayData.Take(heap.Counter)));
                Console.WriteLine(heap.Extract());
            }

            while (true)
            {
                Console.ReadLine();
            }
        }
    }
}
