package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 string = "hello"
    v2 := []int{4294967295, -9223372036854775808, -696}
    var v3 float32 = (((3.14159 * 1.5) - 0.1) - float32(1e10))
    var v4 int8 = 94
    v5 := make(map[uint32]uint32)
    var v6 int64 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        if ((v6 < v6) || false) {
            v8 := GenS1{F0: 1}
            _ = v8
        }
        if (true && (1000 < 962)) {
            if false {
                fmt.Println(fi7)
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(fi7)
                var v10 int64 = v6
                _ = fi9
                _ = v10
            }
        }
        v11 := GenS1{F0: 368}
        _ = fi7
        _ = v11
    }
    var v12 int8 = 42
    v3 += (3.14159 + v3)
    var v13 uint = 2
    if (1 > 32767) {
        fmt.Println(v6)
    }
    v5[782] = 42
    v1 = ""
    var v14 uint16 = 0
    if true {
        if (uint8(v14) > 127) {
            fmt.Println(v3, v6)
        }
        v6 = v6
    } else {
        if false {
            for fi15 := 0; fi15 < 11; fi15++ {
                if true {
                    fmt.Println(v6, v3, v6)
                }
                v16 := GenS1{F0: 42}
                var v17 uint16 = v14
                _ = fi15
                _ = v16
                _ = v17
            }
            switch 32767 {
            case 127:
                var v18 uint16 = (v14 * (^(v14)))
                _ = v18
            }
            switch "" {
            case "foo bar":
                v12 = (v12 + v4)
            }
        }
    }
    var v19 uint32 = 362
    for fi20 := 0; fi20 < 18; fi20++ {
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
        fmt.Println(v14, v1, v13)
        _ = fi20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS1{F0: 102}
        var v23 int32 = -1
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        var v24 float64 = (0.0 * (((1.5 * (0.0 + 1.5)) / 0.1) - 1.5))
        v3 = (v3 - v3)
        _ = v24
    } else {
        v25 := GenS1{F0: 1}
        var v26 int64 = v6
        _ = v25
        _ = v26
    }
    var v27 float32 = v3
    if false {
        var v28 int = 357
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(v3, v3)
            var v30 int32 = -659
            switch (v28 - int(v30)) {
            case -1:
                var v31 int16 = -1
                _ = v31
            case 857:
                v4 = -128
            }
            _ = fi29
            _ = v30
        }
        v4 -= (-(v4))
        _ = v28
    } else {
        v32 := GenS1{F0: 18446744073709551615}
        v19 -= v19
        v3 = (1e10 + -2.25)
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v34 := GenS1{F0: 257}
        v4 = (v4 - (v4 >> (((v13 % ((1) | 1))) & 63)))
        _ = fi33
        _ = v34
    }
    switch v1 {
    case "":
        var v35 int8 = v4
        _ = v35
    case "hello":
        v13 += (v13 + v13)
    }
    v36 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
