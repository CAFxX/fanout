package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint = 42
    var v2 float64 = (1.5 * (1.5 + float64(1.5)))
    var v3 float64 = 0.1
    var v4 uint = 255
    var v5 uint64 = 1
    if (2147483648 >= 551) {
        if (4294967295 >= 1) {
            v2 += v3
        } else {
            fmt.Println(v2)
        }
    } else {
        v6 := GenS1{F0: (^(v1)), F1: -1}
        _ = v6
    }
    var v7 uint8 = 2
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v8 int = 952
        _ = v8
    case "":
        fmt.Println(v4, v2, v3)
    }
    v2 += 1.5
    if (-2.25 <= v2) {
        v3 -= v3
        v9 := []int{32768}
        v9 = append(v9, -270)
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 string = "\n\t"
        _ = fi10
        _ = v11
    }
    var v12 int32 = -2147483648
    v13 := []int{42, 32767, 32767, -482}
    var v14 uint8 = v7
    v13 = append(v13, 1)
    v14 = 104
    for fi15 := 0; fi15 < 4; fi15++ {
        v12 = v12
        _ = fi15
    }
    v16 := GenS2{F0: (v5 - uint64(v1)), F1: 2147483647}
    for fi17 := 0; fi17 < 13; fi17++ {
        v13 = append(v13, fi17)
        _ = fi17
    }
    if ((!(!(-583 != -32768))) && ((true || (true && (-2147483648 == -2147483648))) || ((!("hello" == (("\n\t" + ("hello" + "\n\t")) + ("a" + ((("" + (("\n\t" + "a") + "")) + "") + "foo bar"))))) || false))) {
        fmt.Println(v4, v1, v5)
    } else {
        var v18 uint32 = uint32(v5)
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
