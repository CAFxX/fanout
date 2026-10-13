package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 bool) uint {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{fi7, fi7, fi7}
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    v9 := GenS1{F0: -126, F1: 4294967295}
    var v10 float32 = ((0.0 + 0.0) - (0.0 - 1e10))
    _ = p0
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int = 32768
    var v2 uint16 = 0
    var v3 int32 = -1
    v4 := make(map[uint32]uint32)
    var v5 string = (("foo bar" + ((("foo bar" + "") + "hello") + "")) + "\n\t")
    v6 := make(map[int]string)
    v1 = (v1 & (v1 * (^(v1))))
    if false {
        _ = len(v4)
        v11 := []int{}
        if len(v11) > 0 {
            v11[0] = v1
        }
        _ = v11
    } else {
        switch v1 {
        case 32768:
            _ = len(v4)
        case 296:
            var v12 uint8 = 128
            _ = v12
        }
    }
    v2 -= (v2 * (v2 + v2))
    if true {
        v13 := GenS1{F0: 0, F1: v1}
        var v14 int16 = 1000
        _ = v13
        _ = v14
    }
    var v15 int64 = 9223372036854775807
    v1 -= v1
    v16 := []int{v1, v1, 754, v1}
    var v17 string = ("" + (((((v5 + "\n\t") + v5) + ("a" + (("a" + v5) + ""))) + ("a" + "\n\t")) + ("\n\t" + v5)))
    var v18 int32 = (v3 << ((2) & 63))
    v5 = v5
    var v19 uint8 = 2
    var v20 int16 = 32767
    switch ("a" + ("hello" + ((("" + "") + (v5 + ("" + v5))) + v17))) {
    case "\n\t":
        for fi21 := 0; fi21 < 6; fi21++ {
            v1 = int(v20)
            v15 = 0
            var v22 int16 = (-(v20))
            _ = fi21
            _ = v22
        }
    case "foo bar":
        fmt.Println(v3)
    }
    fmt.Println(v1, v19)
    fmt.Println(v17)
    delete(v4, 2147483648)
    if (v15 >= 1) {
        for fi23 := 0; fi23 < 17; fi23++ {
            switch 127 {
            case 128:
                var v24 int16 = v20
                _ = v24
            case -689:
                v5 = (v5 + v17)
            case 127:
                v25 := v4[0]
                _ = v25
            }
            _ = fi23
        }
        v2 = 65535
        var v26 float32 = ((((((3.14159 * 3.14159) + (1e10 + (0.1 - 1e10))) * 1e10) + 0.1) - float32((3.14159 + 0.1))) + ((0.0 * (((0.0 - float32(-2.25)) - float32((float64(float32(0.1)) - 3.14159))) * (3.14159 + -2.25))) + 1e10))
        _ = v26
    } else {
        if false {
            var v27 uint64 = 4294967295
            _ = v27
        }
        fmt.Println(v15, v2)
        for fi28 := 0; fi28 < 19; fi28++ {
            if len(v16) > 0 {
                v16[0] = fi28
            }
            v29 := GenS1{F0: (v15 >> ((1) & 63)), F1: (fi28 + v1)}
            var v30 float64 = 1.5
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
