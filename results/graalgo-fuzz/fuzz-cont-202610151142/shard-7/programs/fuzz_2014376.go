package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int16; F1 int16; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    var v8 int = 4294967295
    _ = v8
    return 65535
}

func main() {
    var v1 uint = 170
    var v2 uint = (v1 * v1)
    v3 := make(map[int]int)
    var v4 uint = v2
    var v5 uint64 = 2
    v5 = (v5 - 18446744073709551615)
    v9 := GenS1{F0: 0, F1: 42}
    if (19 == -128) {
        var v10 bool = (true || false)
        _ = len(v3)
        _ = v10
    }
    v11 := GenS2{F0: -994, F1: 792, F2: -162}
    v12 := []int{32767}
    var v13 string = ("hello" + "\n\t")
    fmt.Println(v13, v13)
    var v14 int32 = -2147483648
    if true {
        delete(v3, 128)
        var v15 uint32 = 0
        fmt.Println(v13, v15, v14)
        _ = v15
    } else {
        var v16 string = v13
        if true {
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(fi17, fi17)
                _ = fi17
            }
        } else {
            var v18 float64 = 1.5
            v14 += 1
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = v18
        }
        _ = v16
    }
    v4 = v2
    var v19 float32 = 3.14159
    v20 := GenS3{F0: (v13 + ("hello" + "foo bar"))}
    var v21 float64 = 3.14159
    if false {
        var v22 int8 = -110
        v23 := GenS2{F0: -32768, F1: int16(v22), F2: (v14 | v14)}
        fmt.Println(v19, v19, v13)
        _ = v22
        _ = v23
    } else {
        fmt.Println(v5, v1, v5)
        if (42 <= v2) {
            if false {
                fmt.Println(v14, v14)
                var v24 int8 = -128
                v4 = v2
                _ = v24
            } else {
                v25 := GenS4{F0: 65535, F1: 42}
                var v26 bool = true
                _ = v25
                _ = v26
            }
        }
        var v27 int = -652
        _ = v27
    }
    if (1 < 1) {
        delete(v3, -1)
        v28 := v3[1]
        _ = v28
    }
    var v29 int32 = 2147483647
    if false {
        if (1 != (v5 ^ (v5 + v5))) {
            switch v13 {
            case "\n\t":
                if false {
                    var v30 bool = false
                    _ = v30
                }
            case "":
                v31 := GenS2{F0: 32767, F1: 0, F2: -2147483648}
                _ = v31
            case "hello":
                if (42 <= 4294967295) {
                    var v32 int32 = v14
                    _ = v32
                }
            }
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            if false {
                fmt.Println(v1, v29, v21)
                var v33 bool = false
                var v34 uint64 = (v5 << ((2) & 63))
                _ = v33
                _ = v34
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
