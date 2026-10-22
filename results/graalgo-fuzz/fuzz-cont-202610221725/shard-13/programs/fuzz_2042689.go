package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 uint64, p1 int8, p2 int8) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: (-(fi7)), F1: 953, F2: 161}
        fmt.Println(p1, fi7, p2)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 string = "hello"
        _ = fi9
        _ = v10
    }
    fmt.Println(p0, p2)
    if true {
        v11 := GenS1{F0: -396, F1: (p0 % ((p0) | 1)), F2: 42}
        v12 := []int{42, 959, 755, 1}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}
func genFn1(p0 int8, p1 int) int64 {
    for fi13 := 0; fi13 < 9; fi13++ {
        if (!(42 <= 0)) {
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := []int{(^(fi13)), 4294967295, 501}
                v16 := GenS1{F0: p1, F1: 1, F2: 1}
                fmt.Println(fi13, p0, fi13)
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi13
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 float32 = (0.0 - 1e10)
    var v2 int8 = -128
    var v3 uint = 0
    var v4 uint64 = 4294967295
    var v5 int16 = 42
    v6 := make(map[string]uint32)
    v6["\n\t"] = 42
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS1{F0: fi17, F1: (v4 + 2), F2: -620}
        switch "" {
        case "":
            if (true || true) {
                v19 := []int{-9223372036854775808}
                v1 -= v1
                _ = v19
            } else {
                var v20 uint8 = 127
                _ = v20
            }
        case "foo bar":
            if false {
                var v21 bool = (-2.25 == ((1.5 - float64(v1)) * 1e10))
                for fi22 := 0; fi22 < 17; fi22++ {
                    switch (("hello" + "hello") + "\n\t") {
                    case "hello":
                        fmt.Println(fi17, v1, v1)
                    case "":
                        var v23 int32 = 2147483647
                        _ = v23
                    case "foo bar":
                        var v24 uint64 = uint64(v3)
                        _ = v24
                    }
                    _ = fi22
                }
                _ = v21
            } else {
                v25 := GenS1{F0: int(v4), F1: v4, F2: 2147483647}
                _ = v25
            }
        case "hello":
            if true {
                var v26 uint64 = v4
                _ = v26
            }
        }
        switch "\n\t" {
        case "foo bar":
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 float32 = float32(1.5)
                _ = len(v6)
                switch "hello" {
                case "":
                    v4 = 175
                }
                if true {
                    var v29 float32 = v28
                    var v30 uint32 = 732
                    v1 -= ((v29 * 1.5) / v28)
                    _ = v29
                    _ = v30
                } else {
                    var v31 bool = (!(4294967295 <= (v3 + 507)))
                    _ = v31
                }
                _ = fi27
                _ = v28
            }
        }
        _ = fi17
        _ = v18
    }
    var v32 string = "foo bar"
    var v33 uint16 = 0
    var v34 uint32 = 145
    var v35 bool = ((true || false) && true)
    v2 -= v2
    var v36 uint8 = 42
    var v37 uint32 = uint32(v33)
    v38 := GenS1{F0: -9223372036854775808, F1: (^(v4)), F2: -2147483648}
    var v39 uint64 = v4
    switch "\n\t" {
    case "hello":
        var v40 uint16 = 65535
        _ = v40
    case "foo bar":
        v41 := GenS1{F0: -1, F1: 4294967295, F2: -1}
        _ = v41
    case "":
        for fi42 := 0; fi42 < 8; fi42++ {
            fmt.Println(v32)
            delete(v6, ("foo bar" + "foo bar"))
            _ = fi42
        }
    }
    v3 += (v3 + uint(v34))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
