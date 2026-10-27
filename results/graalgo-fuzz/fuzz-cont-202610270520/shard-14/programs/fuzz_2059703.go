package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint32 }
func genFn0() int16 {
    var v5 float32 = 0.1
    var v6 string = "hello"
    v7 := GenS1{F0: -210, F1: 127}
    v8 := []int{}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 972
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    fmt.Println(v1)
    v1 -= 65535
    v9 := []int{-1}
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: ""}
        _ = fi10
        _ = v11
    }
    v1 += v1
    if ((143 == 18446744073709551615) || true) {
        if (!true) {
            v1 = (v1 - 0)
            fmt.Println(v1)
            var v12 int64 = 489
            _ = v12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            switch "a" {
            case "\n\t":
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := GenS1{F0: 2147483647, F1: (fi13 ^ fi13)}
                    var v16 uint8 = 255
                    fmt.Println(v16)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "a":
                if false {
                    v1 += v1
                    if true {
                        v1 -= v1
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v1 += 440
                    }
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            }
            v1 += (v1 | 2)
            _ = fi13
        }
    }
    if (255 > 2) {
        v3[2147483647] = (((("\n\t" + "") + ("hello" + "")) + "hello") + "hello")
        fmt.Println(v1)
    }
    v17 := GenS3{F0: v1}
    v1 += v1
    v18 := v4[-243]
    v18 = (v18 % ((508) | 1))
    var v19 uint = 1
    switch "foo bar" {
    case "":
        v18 -= 0
    case "foo bar":
        v19 = v19
    case "\n\t":
        v20 := GenS4{F0: ("a" > ("hello" + "hello")), F1: 2147483648}
        _ = v20
    }
    if len(v9) > 0 {
        v9[0] = 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
