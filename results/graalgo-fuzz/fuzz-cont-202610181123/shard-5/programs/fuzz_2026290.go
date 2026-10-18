package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0() int32 {
    v7 := GenS1{F0: 1}
    switch "foo bar" {
    case "hello":
        var v8 int32 = 42
        _ = v8
    case "":
        v9 := []int{}
        _ = v9
    default:
        var v10 int16 = -32768
        _ = v10
    }
    var v11 int8 = 2
    v12 := []int{-2147483648, 32768, -415, 42}
    var v13 float32 = -2.25
    _ = v11
    _ = v12
    _ = v13
    _ = v7
    return -826
}

func main() {
    var v1 int = 32767
    var v2 uint8 = 1
    v3 := []string{"", ("\n\t" + "a")}
    v4 := make(map[int]string)
    var v5 uint8 = (v2 << ((4294967295) & 63))
    var v6 uint16 = uint16(v2)
    v6 += 443
    _ = len(v4)
    var v14 uint8 = v5
    switch (("hello" + "a") + "\n\t") {
    case "a":
        v2 += (v2 - (v2 << ((18446744073709551615) & 63)))
    case "":
        if (!(!false)) {
            v2 += 42
            v6 = (v6 / ((v6) | 1))
        }
    case "foo bar":
        v14 = v2
    default:
        var v15 int64 = 32
        _ = v15
    }
    if (!true) {
        var v16 int64 = 494
        var v17 int8 = -95
        var v18 string = ""
        _ = v16
        _ = v17
        _ = v18
    }
    _ = len(v4)
    v6 += (v6 - v6)
    fmt.Println(v2, v1, v6)
    if false {
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 float64 = 1e10
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
            _ = v20
        }
        _ = len(v3)
        v21 := GenS1{F0: 127}
        _ = v21
    } else {
        v4[132] = ""
        if false {
            var v22 uint = 18446744073709551615
            var v23 float64 = 0.1
            var v24 int16 = 42
            _ = v22
            _ = v23
            _ = v24
        }
    }
    if true {
        v6 = 255
        v25 := GenS1{F0: 0}
        v4[v1] = ""
        _ = v25
    } else {
        var v26 int8 = 127
        v26 -= v26
        _ = v26
    }
    switch v1 {
    case 2147483647:
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int = (v1 * -466)
            v29 := GenS1{F0: 100}
            v28 = v28
            _ = fi27
            _ = v28
            _ = v29
        }
    case 276:
        v30 := GenS2{F0: (v6 / (((v6 + 255)) | 1)), F1: 0, F2: 1}
        _ = v30
    case 128:
        v2 -= (v14 >> ((2) & 63))
    default:
        v14 -= (v2 & (v2 >> ((2) & 63)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
