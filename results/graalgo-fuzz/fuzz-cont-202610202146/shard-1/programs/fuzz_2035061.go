package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint }

func main() {
    var v1 string = "a"
    var v2 bool = false
    var v3 string = v1
    var v4 uint16 = 2
    var v5 uint8 = 129
    var v6 int = 32768
    v7 := GenS1{F0: (v4 & v4), F1: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        switch v1 {
        case "\n\t":
            v9 := GenS2{F0: -291}
            _ = v9
        case "a":
            v6 -= 32768
        case "":
            var v10 uint = 561
            _ = v10
        }
        var v11 uint64 = 697
        _ = fi8
        _ = v11
    }
    switch ("a" + ((((((((v3 + v1) + "") + v3) + v1) + "hello") + v3) + (("hello" + ((v1 + (v3 + v1)) + v1)) + v3)) + "\n\t")) {
    case "":
        v5 -= (v5 / ((v5) | 1))
    case "hello":
        fmt.Println(v6)
    }
    v12 := GenS2{F0: 9223372036854775807}
    if v2 {
        var v13 bool = (2147483648 >= 2147483647)
        v5 += (v5 << ((255) & 63))
        _ = v13
    }
    v14 := []int{int(v5), v6, v6}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if v2 {
        if (65535 > 1) {
            if true {
                fmt.Println(v1)
                switch (v6 + v6) {
                case 32767:
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS3{F0: -428}
                        fmt.Println(v1, fi15, v3)
                        v2 = ((72 != -128) || v2)
                        _ = fi15
                        _ = v16
                    }
                }
            } else {
                v17 := GenS1{F0: v4, F1: 1}
                _ = v17
            }
        }
    } else {
        v18 := GenS4{F0: 4294967295, F1: v6, F2: 0}
        v19 := GenS3{F0: 9223372036854775807}
        _ = v18
        _ = v19
    }
    v1 = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
