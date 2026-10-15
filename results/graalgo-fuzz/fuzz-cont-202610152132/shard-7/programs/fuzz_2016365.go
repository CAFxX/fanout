package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int32; F2 bool }

func main() {
    var v1 uint8 = 32
    var v2 uint32 = 457
    v3 := make(map[string]int)
    var v4 string = "\n\t"
    var v5 uint16 = 929
    var v6 uint = 830
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int = 127
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{(v8 ^ fi7), (^(v8))}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 int = -9223372036854775808
    v12 := GenS1{F0: v4}
    switch v11 {
    case 326:
        if false {
            var v13 string = (v4 + v4)
            _ = v13
        } else {
            var v14 uint = 65535
            v15 := v3["foo bar"]
            _ = v14
            _ = v15
        }
    }
    var v16 uint8 = v1
    if false {
        v17 := []int{v11, v11, 519}
        v4 = v4
        v18 := GenS1{F0: "hello"}
        _ = v17
        _ = v18
    }
    switch v11 {
    case -276:
        if false {
            switch (v11 >> ((v6) & 63)) {
            case 9223372036854775807:
                v16 += 128
            default:
                var v19 int = (-(v11))
                _ = v19
            }
            if (!(v5 > v5)) {
                v6 -= v6
                fmt.Println(v5)
                if (316 > -32768) {
                    v20 := GenS1{F0: ((v4 + ("" + (v4 + v4))) + "")}
                    if (4294967295 <= v2) {
                        v21 := []int{v11}
                        _ = v21
                    }
                    _ = v20
                } else {
                    v22 := GenS2{F0: -160, F1: 108}
                    if true {
                        v4 = "\n\t"
                    }
                    v3[v4] = 88
                    _ = v22
                }
            }
        }
    case 941:
        v23 := GenS2{F0: -1, F1: 127}
        _ = v23
    case -2147483648:
        v4 = (v4 + (v4 + v4))
    }
    v24 := GenS2{F0: -1, F1: 78}
    var v25 uint64 = 755
    v26 := GenS2{F0: 0, F1: 100}
    fmt.Println(v25, v6, v6)
    v27 := GenS3{F0: 32767}
    switch ((v4 + v4) + (v4 + v4)) {
    case "a":
        v28 := GenS4{F0: v16}
        _ = v28
    case "hello":
        for fi29 := 0; fi29 < 14; fi29++ {
            if true {
                var v30 float32 = (0.0 + 0.0)
                v1 = (v16 << ((647) & 63))
                _ = v30
            }
            v2 += 680
            _ = fi29
        }
    default:
        if (!(true || ((v25 % (((v25 / (((v25 + 42)) | 1))) | 1)) <= uint64(v5)))) {
            v31 := GenS1{F0: ("foo bar" + ("a" + v4))}
            v11 = v11
            if (42 != -45) {
                v32 := GenS5{F0: (v5 * 255), F1: -1, F2: ((false || true) || true)}
                _ = v32
            }
            _ = v31
        } else {
            var v33 uint32 = v2
            v34 := GenS5{F0: v5, F1: -511, F2: false}
            v35 := []int{474}
            _ = v33
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
