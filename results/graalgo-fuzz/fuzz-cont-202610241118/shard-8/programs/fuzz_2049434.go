package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1000
    v3 := make(map[string]string)
    var v4 uint = 0
    var v5 uint64 = (v1 * (v1 & v1))
    v5 = v1
    fmt.Println(v4, v2)
    if false {
        var v6 int32 = 1
        fmt.Println(v4, v6, v4)
        fmt.Println(v6, v5)
        _ = v6
    }
    v7 := []int{172}
    fmt.Println(v4, v1, v1)
    if true {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint8 = 1
            for fi10 := 0; fi10 < 18; fi10++ {
                switch ("a" + "a") {
                case "hello":
                    if true {
                        v11 := v3["a"]
                        v4 += v4
                        _ = v11
                    } else {
                        v2 = -1
                    }
                }
                v5 += 901
                switch ("\n\t" + "\n\t") {
                case "a":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                v9 = 0
                _ = len(v3)
                _ = fi10
            }
            if true {
                v5 += v5
            } else {
                v7 = append(v7, fi8)
                v9 -= v9
            }
            _ = fi8
            _ = v9
        }
        v12 := GenS1{F0: 1}
        v5 += v5
        _ = v12
    }
    var v13 int8 = 0
    var v14 uint64 = 357
    v4 -= v4
    var v15 uint = (v4 * 2)
    fmt.Println(v13, v1)
    fmt.Println(v14)
    var v16 int32 = 0
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v14)
        _ = fi17
    }
    var v18 int = -503
    switch int(v15) {
    case -9223372036854775808:
        v15 += (^(v15))
    case 128:
        v7 = append(v7, v18)
    case 4294967295:
        var v19 int8 = (v13 % ((1) | 1))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
