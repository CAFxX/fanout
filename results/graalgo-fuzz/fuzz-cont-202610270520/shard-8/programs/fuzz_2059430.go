package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint8; F2 string }
type GenS5 struct { F0 int; F1 int64; F2 int32 }

func main() {
    var v1 int8 = 1
    v2 := []int{128}
    var v3 int = -417
    v4 := make(map[string]int)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 := GenS1{F0: 19, F1: 248}
    v3 = (v3 >> ((1) & 63))
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            if len(v2) > 0 {
                v2[0] = 1
            }
            var v8 uint64 = 1
            v1 = v1
            _ = v8
        } else {
            var v9 uint16 = 65535
            for fi10 := 0; fi10 < 11; fi10++ {
                if false {
                    fmt.Println(fi10, v1, v9)
                    if (65535 >= 1) {
                        v4["a"] = (-(fi7))
                    }
                }
                _ = fi10
            }
            _ = v9
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
        _ = fi7
    }
    v6 -= v6
    v6 += (^(v6))
    delete(v4, ("a" + "foo bar"))
    if (2147483647 == 4294967295) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 float32 = ((float32((((float64((-2.25 * -2.25)) / 3.14159) + 3.14159) / 1.0)) / (0.1 + 3.14159)) + 0.1)
            _ = fi12
            _ = v13
        }
        if true {
            var v14 uint = 919
            if (0.1 == 1.5) {
                v15 := GenS1{F0: (v6 << (((v14 + v14)) & 63)), F1: v6}
                _ = v15
            }
            v16 := GenS1{F0: 42, F1: v6}
            _ = v14
            _ = v16
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 10; fi18++ {
                    fmt.Println(fi17, v1, v1)
                    _ = fi18
                }
                v3 -= (fi17 | v3)
                _ = fi17
            }
        }
    } else {
        v19 := GenS2{F0: "hello"}
        if len(v2) > 0 {
            v2[0] = (v3 / ((2147483647) | 1))
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 uint32 = 278
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    delete(v4, ("a" + "a"))
    v3 = v3
    var v22 uint16 = 277
    v23 := GenS1{F0: (v6 & (v6 << ((18446744073709551615) & 63))), F1: v6}
    if (!true) {
        v24 := GenS3{F0: 0, F1: v6}
        _ = v24
    } else {
        v25 := GenS4{F0: -9223372036854775808, F1: v6, F2: "a"}
        _ = v25
    }
    if true {
        if true {
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 string = "a"
                v28 := GenS5{F0: (^(v3)), F1: -1, F2: 2147483647}
                v27 = (("" + ("\n\t" + "")) + v27)
                _ = fi26
                _ = v27
                _ = v28
            }
            switch (v3 >> ((4294967295) & 63)) {
            case 32767:
                var v29 uint64 = 42
                _ = v29
            case -1:
                var v30 uint16 = 255
                _ = v30
            }
        }
        if len(v2) > 0 {
            v2[0] = v3
        }
    } else {
        switch "a" {
        case "":
            switch "foo bar" {
            case "":
                if len(v2) > 0 {
                    v2[0] = 128
                }
            }
        case "a":
            var v31 uint32 = 2
            _ = v31
        }
    }
    var v32 int = (-(v3))
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(v1, v22)
        fmt.Println(v3)
        var v34 string = ("a" + "\n\t")
        _ = fi33
        _ = v34
    }
    fmt.Println(v6, v22, v22)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
