package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint; F2 int16 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        switch "" {
        case "hello":
            switch ("hello" + ("hello" + "")) {
            case "a":
                var v5 uint8 = 1
                _ = v5
            default:
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(fi4)
                case "\n\t":
                    var v6 int16 = 1
                    _ = v6
                case "hello":
                    var v7 string = "hello"
                    _ = v7
                }
            }
        }
        var v8 string = ("foo bar" + "a")
        var v9 uint = 65535
        _ = fi4
        _ = v8
        _ = v9
    }
    var v10 int8 = 1
    v11 := []int{}
    v10 = v10
    v10 = v10
    _ = v10
    _ = v11
    return 104
}

func main() {
    var v1 int16 = 42
    var v2 float32 = 3.14159
    var v3 int16 = v1
    fmt.Println(v2, v3)
    var v12 uint64 = 0
    switch 582 {
    case -1:
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                if ((!(true || (315 != 4294967295))) || false) {
                    v15 := GenS1{F0: 18446744073709551615}
                    _ = v15
                }
                var v16 bool = true
                v3 = (^(v3))
                _ = fi14
                _ = v16
            }
            _ = fi13
        }
    default:
        for fi17 := 0; fi17 < 19; fi17++ {
            if true {
                if (!(false || true)) {
                    v18 := []int{(fi17 % (((fi17 | fi17)) | 1)), (fi17 | 2147483647), -9223372036854775808}
                    _ = v18
                } else {
                    v1 -= v1
                    v19 := []int{(fi17 & (fi17 & fi17)), (fi17 & (fi17 ^ fi17))}
                    _ = v19
                }
                v20 := []int{fi17, 678, fi17}
                for fi21 := 0; fi21 < 17; fi21++ {
                    fmt.Println(v3, fi21, fi17)
                    var v22 uint64 = 18446744073709551615
                    v3 = 42
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                var v23 bool = (fi17 >= (fi17 << ((168) & 63)))
                _ = v23
            }
            _ = fi17
        }
    }
    var v24 uint8 = genFn0()
    fmt.Println(v2)
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := []int{(fi25 - -53), fi25, fi25, fi25}
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 65535}
    var v28 string = "a"
    if ((42 != -128) && (!((3.14159 - float32((3.14159 / 1.5))) == ((-2.25 * 0.1) * 1e10)))) {
        v29 := GenS1{F0: 18446744073709551615}
        _ = v29
    } else {
        var v30 uint = uint(v12)
        v28 = v28
        var v31 bool = false
        _ = v30
        _ = v31
    }
    v32 := []int{-1, -9223372036854775808, 128, 322}
    fmt.Println(v28, v24, v1)
    var v33 float64 = 0.0
    if len(v32) > 0 {
        v32[0] = 0
    }
    if len(v32) > 0 {
        v32[0] = 1
    }
    v33 = v33
    if true {
        v34 := GenS2{F0: (v28 + ""), F1: 1, F2: (v3 >> ((uint(v12)) & 63))}
        var v35 uint = 18446744073709551615
        v28 = v28
        _ = v34
        _ = v35
    }
    fmt.Println(v28, v33, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
}
