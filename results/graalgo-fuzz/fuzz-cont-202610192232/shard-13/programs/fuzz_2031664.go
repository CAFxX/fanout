package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32 }
func genFn0() string {
    switch ("\n\t" + "\n\t") {
    case "hello":
        v5 := GenS1{F0: -358}
        _ = v5
    case "":
        var v6 int = -2147483648
        _ = v6
    }
    var v7 uint8 = 183
    var v8 int = 2147483647
    var v9 float64 = float64(-2.25)
    var v10 uint32 = 429
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (("foo bar" + (("a" + "foo bar") + "a")) + "a")
}

func main() {
    var v1 bool = false
    var v2 int16 = -390
    var v3 int64 = -9223372036854775808
    var v4 uint = 0
    v4 += (v4 ^ v4)
    var v11 uint16 = 0
    if (("" + "\n\t") != "\n\t") {
        var v12 uint32 = 915
        var v13 uint64 = 937
        _ = v12
        _ = v13
    }
    v1 = (-128 >= 100)
    if (false && true) {
        if v1 {
            v14 := []int{-683, -1, 744, 676}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v14
        }
    }
    var v15 float32 = 3.14159
    var v16 uint8 = 78
    switch ("a" + "\n\t") {
    case "":
        v11 -= v11
    case "\n\t":
        fmt.Println(v11, v4)
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            var v17 float64 = ((1e10 - float64(float32(1.5))) / 1.5)
            _ = v17
        case "":
            var v18 int16 = v2
            _ = v18
        }
    default:
        v19 := GenS1{F0: (v3 << ((686) & 63))}
        _ = v19
    }
    var v20 int8 = 0
    v21 := GenS2{F0: v11, F1: 125, F2: (v20 << ((65535) & 63))}
    switch int(v3) {
    case 32768:
        var v22 int8 = (v20 * 1)
        _ = v22
    default:
        v23 := GenS1{F0: v3}
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS1{F0: -9223372036854775808}
        var v26 uint16 = 2
        for fi27 := 0; fi27 < 19; fi27++ {
            v1 = false
            _ = fi27
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v28 string = ("hello" + ("" + "a"))
    var v29 int16 = (v2 / (((v2 ^ -32768)) | 1))
    v30 := GenS3{F0: (v2 | (v2 ^ v2)), F1: v3}
    for fi31 := 0; fi31 < 10; fi31++ {
        v15 = v15
        var v32 int8 = 127
        v33 := GenS4{F0: int32(v2)}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
