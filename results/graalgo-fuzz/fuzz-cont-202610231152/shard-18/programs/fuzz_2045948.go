package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 int16 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 string, p1 uint32) int8 {
    var v6 int = 725
    for fi7 := 0; fi7 < 3; fi7++ {
        switch ((("\n\t" + "") + "\n\t") + "a") {
        case "":
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 uint64 = 98
                switch (fi7 << ((uint(v9)) & 63)) {
                case 127:
                    if (!false) {
                        v10 := GenS1{F0: 100}
                        var v11 uint64 = (v9 & 0)
                        var v12 uint32 = (p1 - 514)
                        var v13 uint32 = v12
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                _ = fi8
                _ = v9
            }
        case "a":
            var v14 uint64 = 4294967295
            _ = v14
        case "\n\t":
            var v15 uint8 = 7
            _ = v15
        }
        _ = fi7
    }
    v16 := []int{42, (v6 * 148), -2147483648, (-(v6))}
    for fi17 := 0; fi17 < 20; fi17++ {
        v6 = (v6 << ((18446744073709551615) & 63))
        for fi18 := 0; fi18 < 7; fi18++ {
            if true {
                var v19 float32 = 0.1
                var v20 bool = false
                switch v6 {
                case -1:
                    fmt.Println(p1)
                    var v21 float64 = 1.5
                    p0 = ("" + ((p0 + p0) + p0))
                    _ = v21
                case 127:
                    var v22 float32 = v19
                    _ = v22
                case 32768:
                    switch int(p1) {
                    case -1:
                        v19 = 0.0
                    case -274:
                        v23 := GenS1{F0: 127}
                        _ = v23
                    case 1:
                        p0 = ("" + p0)
                    }
                }
                _ = v19
                _ = v20
            } else {
                v24 := GenS1{F0: 43}
                v6 = v6
                _ = v24
            }
            v6 += fi18
            var v25 uint16 = 881
            _ = fi18
            _ = v25
        }
        v16 = append(v16, (-(v6)))
        _ = fi17
    }
    v26 := GenS1{F0: 60}
    _ = p0
    _ = p1
    _ = v16
    _ = v26
    _ = v6
    return 1
}
func genFn1(p0 uint32, p1 bool) int {
    var v27 int32 = 241
    v28 := GenS1{F0: 53}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return 2147483647
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{2147483648, 2147483648}
    v3 := make(map[int]bool)
    var v4 bool = false
    var v5 int64 = (v1 | (v1 * -1))
    var v29 uint16 = 32767
    if v4 {
        var v30 int16 = 1
        fmt.Println(v5)
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v29 = 42
        fmt.Println(v5)
        v1 = (v5 & (^(v5)))
        _ = fi31
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 uint16 = v29
        var v34 int8 = genFn0("", 42)
        var v35 uint8 = 42
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v5 -= v1
    v36 := GenS1{F0: genFn0((("hello" + ("\n\t" + "hello")) + "a"), 42)}
    if (true || (!(!(v4 && v4)))) {
        v37 := GenS1{F0: -83}
        for fi38 := 0; fi38 < 5; fi38++ {
            fmt.Println(fi38)
            v39 := GenS1{F0: 42}
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        v40 := GenS2{F0: -1, F1: ("hello" + ((("a" + "hello") + "hello") + ("" + "foo bar"))), F2: 32767}
        var v41 float64 = 1.5
        _ = v40
        _ = v41
    }
    switch ("" + (("foo bar" + "") + "")) {
    case "a":
        v1 = (v5 & (-(v1)))
    case "foo bar":
        switch "a" {
        case "a":
            fmt.Println(v5, v29)
        case "foo bar":
            v42 := GenS3{F0: false, F1: 902}
            _ = v42
        case "hello":
            if (0 >= 828) {
                v5 = v1
                v43 := GenS4{F0: 737, F1: (v29 | (v29 % (((v29 | 552)) | 1)))}
                _ = v43
            }
        default:
            var v44 bool = ((((("\n\t" + "hello") + "foo bar") + (("hello" + "hello") + "hello")) + "a") <= (("a" + (("" + ((("\n\t" + "") + "a") + "")) + (("a" + "\n\t") + "hello"))) + ""))
            _ = v44
        }
    case "\n\t":
        v4 = (v4 || v4)
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if v4 {
        v45 := v3[2147483647]
        _ = v45
    }
    fmt.Println(v5, v4, v29)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
