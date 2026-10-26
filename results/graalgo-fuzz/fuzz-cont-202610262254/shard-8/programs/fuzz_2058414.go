package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint16, p1 bool) uint16 {
    if true {
        v6 := []int{4294967295, 42, 9223372036854775807, -2147483648}
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int32 = 687
            var v9 uint16 = p0
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        v10 := GenS1{F0: "hello", F1: 1}
        v11 := []int{}
        v12 := GenS1{F0: "a", F1: 2}
        _ = v10
        _ = v11
        _ = v12
    }
    switch ("\n\t" + ("" + (("hello" + "\n\t") + "a"))) {
    case "hello":
        if (2147483647 > 2147483647) {
            switch -426 {
            case 1:
                fmt.Println(p0)
            case 32767:
                fmt.Println(p1, p0)
            default:
                var v13 float64 = 3.14159
                _ = v13
            }
            if (p1 || true) {
                v14 := []int{2147483647}
                _ = v14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                p1 = true
                _ = fi15
            }
        } else {
            if true {
                p1 = false
            }
        }
    case "a":
        switch "hello" {
        case "hello":
            v16 := []int{32768, -28, -2147483648, -1}
            _ = v16
        case "foo bar":
            if (("a" + "\n\t") < ("" + ("a" + "a"))) {
                v17 := GenS2{F0: 128}
                if p1 {
                    v18 := []int{-226}
                    v19 := GenS3{F0: 0, F1: 42, F2: 255}
                    _ = v18
                    _ = v19
                }
                v20 := GenS2{F0: 134}
                _ = v17
                _ = v20
            }
        case "a":
            v21 := []int{}
            _ = v21
        }
    case "foo bar":
        var v22 int8 = 109
        _ = v22
    }
    var v23 int16 = -1
    var v24 string = ("foo bar" + ("" + "a"))
    for fi25 := 0; fi25 < 7; fi25++ {
        fmt.Println(v23, v23)
        p1 = (!p1)
        var v26 int = (-(fi25))
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return p0
}

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 uint32 = 65535
    var v4 uint16 = 255
    var v5 uint8 = 42
    var v27 uint = 4294967295
    var v28 int16 = -24
    var v29 bool = false
    v3 = v3
    v5 -= v5
    switch ("a" + ("" + "a")) {
    case "a":
        if v29 {
            v27 -= v27
        } else {
            var v30 int32 = -2147483648
            _ = v30
        }
    case "hello":
        v31 := GenS1{F0: "hello", F1: 4294967295}
        _ = v31
    }
    var v32 int8 = 127
    v4 -= (^(v4))
    if ((v28 % ((v28) | 1)) == int16(v27)) {
        for fi33 := 0; fi33 < 8; fi33++ {
            if (840 == -2147483648) {
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := GenS1{F0: "hello", F1: 42}
                    if (uint16(v27) != (^(v4))) {
                        v36 := GenS4{F0: (v3 >> (((v27 ^ v27)) & 63)), F1: (-(fi34))}
                        if v1 {
                            v37 := GenS5{F0: v4, F1: -2147483648}
                            _ = v37
                        }
                        fmt.Println(v32)
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
                var v38 int = (fi33 << (((v27 % ((18446744073709551615) | 1))) & 63))
                v1 = (!(!(2 != 948)))
                _ = v38
            }
            switch (^(fi33)) {
            case 657:
                v39 := []int{}
                _ = v39
            case 801:
                v40 := []int{}
                _ = v40
            case 350:
                fmt.Println(v1, v28)
            default:
                v41 := v2["hello"]
                _ = v41
            }
            var v42 uint = (v27 / ((v27) | 1))
            _ = fi33
            _ = v42
        }
        v28 = (v28 >> ((v27) & 63))
        if false {
            v43 := []int{}
            _ = v43
        }
    }
    var v44 uint64 = 556
    for fi45 := 0; fi45 < 12; fi45++ {
        if true {
            for fi46 := 0; fi46 < 12; fi46++ {
                var v47 int = fi46
                v48 := GenS4{F0: (^(v3)), F1: (fi45 - int(v4))}
                _ = fi46
                _ = v47
                _ = v48
            }
            var v49 float32 = (3.14159 * 3.14159)
            v3 -= (^(v3))
            _ = v49
        }
        _ = fi45
    }
    v50 := v2["hello"]
    for fi51 := 0; fi51 < 2; fi51++ {
        v27 = v27
        _ = fi51
    }
    switch 40 {
    case 260:
        v32 = 0
    case 42:
        _ = len(v2)
    }
    v5 = uint8(v3)
    if v29 {
        v52 := GenS5{F0: genFn0(uint16(v27), v1), F1: 649}
        v4 -= (v4 - uint16(v44))
        _ = v52
    } else {
        fmt.Println(v27, v29, v3)
    }
    v2["hello"] = (v50 + (v50 - uint32(v4)))
    fmt.Println(v1, v3, v44)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v5
    _ = v50
}
