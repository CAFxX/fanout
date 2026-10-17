package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int64; F2 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 int) int16 {
    var v7 int16 = -32768
    var v8 int64 = 645
    switch (p1 | (p1 << ((0) & 63))) {
    case 0:
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int16 = (v7 & 0)
            fmt.Println(p0, v8, p1)
            var v11 int64 = v8
            _ = fi9
            _ = v10
            _ = v11
        }
    case 1:
        p0 = 2
    case -9223372036854775808:
        v12 := []int{(p1 - -667), -9223372036854775808, p1}
        _ = v12
    }
    v13 := []int{321, 573}
    switch (p1 | p1) {
    case -1:
        var v14 uint32 = 2
        _ = v14
    case -929:
        var v15 int16 = (v7 | v7)
        _ = v15
    case -886:
        v16 := GenS1{F0: 4294967295}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{2, 99}
    var v3 int16 = 0
    v4 := []uint32{}
    v5 := []string{"a", ("hello" + (((("hello" + (("hello" + "") + "")) + "a") + ((("hello" + "hello") + "\n\t") + "\n\t")) + "a")), ("a" + "hello")}
    var v6 float64 = (float64(0.0) * 0.1)
    v1 += 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += (v1 ^ v1)
    if (2147483647 == 1) {
        if len(v2) > 0 {
            v2[0] = 856
        }
        v6 += v6
        switch "\n\t" {
        case "hello":
            var v17 int8 = 42
            _ = v17
        case "\n\t":
            var v18 int = 128
            _ = v18
        case "":
            var v19 int32 = -2147483648
            _ = v19
        default:
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "":
                for fi20 := 0; fi20 < 12; fi20++ {
                    fmt.Println(v3, v1)
                    _ = fi20
                }
            case "\n\t":
                v5 = append(v5, ("a" + ("\n\t" + ("foo bar" + ("" + "a")))))
            case "hello":
                switch 658 {
                case 0:
                    var v21 int16 = (v3 * -32768)
                    _ = v21
                }
            }
        }
    }
    var v22 float64 = (v6 / float64(3.14159))
    var v23 uint16 = 130
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v24 int8 = 42
    switch "" {
    case "":
        if (-9223372036854775808 >= int64(v1)) {
            if true {
                v25 := GenS2{F0: 4294967295, F1: 770, F2: (v3 | genFn0(320, -309))}
                var v26 float32 = 1e10
                _ = v25
                _ = v26
            } else {
                if ((-1 < 868) && true) {
                    var v27 uint32 = 2147483647
                    _ = v27
                }
                v23 += 2
            }
            for fi28 := 0; fi28 < 14; fi28++ {
                v6 -= 0.1
                fmt.Println(fi28, v3, fi28)
                _ = fi28
            }
            v24 = v24
        }
    default:
        if (true || true) {
            var v29 uint32 = 0
            v30 := GenS3{F0: 485}
            var v31 int8 = -128
            _ = v29
            _ = v30
            _ = v31
        }
    }
    fmt.Println(v1)
    fmt.Println(v6, v22)
    fmt.Println(v23, v1)
    var v32 float32 = 1.5
    v23 += (^(v23))
    for fi33 := 0; fi33 < 8; fi33++ {
        if true {
            var v34 uint32 = 2147483648
            v6 = v22
            v35 := GenS3{F0: v3}
            _ = v34
            _ = v35
        }
        var v36 uint = 0
        if len(v2) > 0 {
            v2[0] = 2
        }
        _ = fi33
        _ = v36
    }
    var v37 string = "hello"
    switch 128 {
    case 127:
        var v38 uint64 = v1
        _ = v38
    case 9223372036854775807:
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                if (-64 != -1) {
                    fmt.Println(v3)
                } else {
                    var v41 float32 = (1.5 - v32)
                    _ = len(v5)
                    _ = v41
                }
                _ = fi40
            }
            switch "a" {
            case "a":
                switch "hello" {
                case "a":
                    v5 = append(v5, "\n\t")
                case "hello":
                    v6 = -2.25
                case "":
                    var v42 uint8 = 163
                    _ = v42
                default:
                    var v43 int = fi39
                    _ = v43
                }
            }
            v22 = v6
            _ = fi39
        }
    case 63:
        v37 = "hello"
    }
    v6 = v22
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
