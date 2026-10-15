package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() bool {
    var v5 uint16 = 32767
    v6 := []int{4294967295, -9223372036854775808, 877, -348}
    _ = v5
    _ = v6
    return (true || false)
}
func genFn1() int64 {
    var v7 bool = true
    v8 := GenS1{F0: 729}
    v7 = (1000 == -32768)
    if (42 != 831) {
        v9 := []int{32767}
        v7 = v7
        fmt.Println(v7)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(v7, fi11)
                var v12 uint = 1
                if v7 {
                    var v13 int64 = 30
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            v14 := []int{}
            _ = len(v14)
            _ = fi10
            _ = v14
        }
    }
    v7 = (v7 && (true || (-1 < 42)))
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int = -158
    v3 := make(map[string]int)
    v4 := []int{(v2 - v2), 1, v2, v2}
    fmt.Println(v2)
    if true {
        v2 += (^(v2))
    }
    v1[("foo bar" + ("foo bar" + "a"))] = ("hello" + "hello")
    v2 = -2147483648
    var v15 uint = 2
    var v16 bool = true
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v2, v2, fi17)
        v16 = (v16 && false)
        v15 = (v15 << (((v15 * v15)) & 63))
        _ = fi17
    }
    fmt.Println(v16, v15, v2)
    var v18 string = ("\n\t" + "\n\t")
    fmt.Println(v18, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
