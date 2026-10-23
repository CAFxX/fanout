package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int32; F2 int32 }

func main() {
    var v1 int16 = -1
    var v2 uint32 = 2
    var v3 float64 = 1.5
    v4 := make(map[int]uint32)
    var v5 int64 = 955
    var v6 int32 = 152
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint8 = 215
        _ = len(v4)
        var v9 int8 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 -= v6
    for fi10 := 0; fi10 < 19; fi10++ {
        if true {
            if false {
                fmt.Println(v5)
                fmt.Println(v1, v3)
                var v11 int = -406
                _ = v11
            }
            fmt.Println(fi10, v6, v1)
        } else {
            var v12 string = "\n\t"
            delete(v4, (fi10 * fi10))
            var v13 bool = false
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            v2 += (v2 * 4294967295)
            _ = fi14
        }
        _ = fi10
    }
    var v15 bool = true
    v15 = (65535 >= 114)
    switch (("\n\t" + ("foo bar" + ("a" + ("" + "hello")))) + "a") {
    case "a":
        v16 := GenS1{F0: v5}
        _ = v16
    case "hello":
        v17 := GenS2{F0: 1, F1: v6, F2: (^(v6))}
        _ = v17
    }
    if (!(!((((v5 | (v5 * v5)) < (v5 >> ((18446744073709551615) & 63))) || (true || ((42 == 65535) && (1 < 71)))) || (!false)))) {
        var v18 int = 4294967295
        v19 := []int{v18, (v18 % (((v18 & v18)) | 1)), (v18 >> ((701) & 63))}
        var v20 uint64 = 104
        _ = v18
        _ = v19
        _ = v20
    }
    v1 -= v1
    var v21 int = 38
    var v22 int16 = 32767
    if false {
        if true {
            v23 := GenS1{F0: (-(v5))}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                v15 = v15
                _ = fi24
            }
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            fmt.Println(v15, v5, fi25)
            switch "a" {
            case "foo bar":
                if (!v15) {
                    v21 = 9223372036854775807
                    var v26 float64 = 3.14159
                    _ = v26
                }
            case "hello":
                var v27 uint32 = v2
                _ = v27
            default:
                var v28 int32 = int32(v21)
                _ = v28
            }
            var v29 int64 = (v5 * (v5 >> ((4294967295) & 63)))
            _ = fi25
            _ = v29
        }
        v6 = v6
    } else {
        var v30 uint8 = 127
        _ = v30
    }
    var v31 uint = 725
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 string = "hello"
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
