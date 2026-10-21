package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = v1
    var v3 int8 = -1
    var v4 bool = true
    var v5 bool = (true && (((("a" + ("a" + (((("" + ("a" + "\n\t")) + "a") + "a") + ("a" + "hello")))) > "a") || false) && v4))
    if v4 {
        v5 = v4
        v6 := GenS1{F0: (v3 % ((100) | 1))}
        for fi7 := 0; fi7 < 9; fi7++ {
            v4 = false
            v8 := GenS1{F0: v3}
            fmt.Println(v2)
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    v4 = v4
    v3 = (v3 / ((v3) | 1))
    if (!((false || true) && (((("hello" + (((("\n\t" + "hello") + "a") + "a") + ("foo bar" + "\n\t"))) >= ("\n\t" + ("a" + "hello"))) || v4) && false))) {
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                switch fi10 {
                case 0:
                    var v11 uint32 = 863
                    _ = v11
                case -100:
                    v12 := []int{fi9}
                    _ = v12
                case -853:
                    switch -9223372036854775808 {
                    case 32767:
                        v2 = (v1 | v1)
                    default:
                        var v13 uint64 = 42
                        _ = v13
                    }
                }
                _ = fi10
            }
            v14 := []int{42, fi9, 32768}
            v15 := GenS1{F0: (v3 | (v3 >> ((150) & 63)))}
            _ = fi9
            _ = v14
            _ = v15
        }
    } else {
        v4 = (255 > 326)
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v5)
    var v18 uint32 = 0
    v19 := GenS1{F0: (v3 & (v3 - v3))}
    var v20 uint16 = 255
    v21 := GenS2{F0: (v18 << ((255) & 63)), F1: (-(v3)), F2: (v20 & (v20 >> ((4294967295) & 63)))}
    if (!(v2 >= (v2 % (((^(v2))) | 1)))) {
        var v22 uint64 = 2
        fmt.Println(v22, v22, v1)
        var v23 uint16 = 0
        _ = v22
        _ = v23
    } else {
        v5 = true
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: -13}
        _ = fi24
        _ = v25
    }
    var v26 uint32 = (^(v18))
    fmt.Println(v18, v3)
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := []int{335, int(v1), fi27, (fi27 | -738)}
        v28 = append(v28, -9223372036854775808)
        var v29 int8 = v3
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int32 = 2147483647
    var v31 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
