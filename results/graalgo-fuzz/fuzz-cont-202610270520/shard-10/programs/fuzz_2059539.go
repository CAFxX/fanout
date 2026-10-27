package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint64; F2 bool }
func genFn0() int64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{}
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 float64 = 3.14159
            var v9 int64 = 42
            if true {
                v9 += (v9 / ((v9) | 1))
                var v10 int = 0
                var v11 string = "foo bar"
                _ = v10
                _ = v11
            } else {
                v12 := GenS1{F0: v9, F1: -1}
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS2{F0: 1}
        var v15 bool = true
        if true {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint64 = 2
                v18 := GenS1{F0: -9223372036854775808, F1: 573}
                switch ((("a" + "a") + ("foo bar" + "")) + ("a" + ("a" + "hello"))) {
                case "\n\t":
                    fmt.Println(fi13, v17)
                case "foo bar":
                    if v15 {
                        fmt.Println(fi16)
                        var v19 uint32 = 0
                        _ = v19
                    }
                case "hello":
                    v15 = (766 != 323)
                }
                var v20 uint64 = (v17 - (v17 | 347))
                _ = fi16
                _ = v17
                _ = v18
                _ = v20
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 9223372036854775807 {
    case -613:
        if true {
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 bool = (288 <= int16(fi21))
                _ = fi21
                _ = v22
            }
            if true {
                var v23 uint64 = 725
                v24 := GenS3{F0: -9223372036854775808, F1: 113}
                var v25 float64 = -2.25
                _ = v23
                _ = v24
                _ = v25
            } else {
                if (42 <= 0) {
                    v26 := []int{-1, -748, 633}
                    var v27 uint = 1
                    _ = v26
                    _ = v27
                }
            }
        }
    case 128:
        var v28 uint16 = 206
        _ = v28
    }
    var v29 float64 = float64(((0.1 + (1e10 * 0.0)) + (float32((-2.25 / 1.0)) * (1.5 - 1.5))))
    _ = v29
    return -1
}
func genFn1(p0 uint8, p1 uint16, p2 int8) uint16 {
    p1 = (p1 % ((p1) | 1))
    v30 := GenS1{F0: 0, F1: 881}
    if true {
        var v31 int = 42
        _ = v31
    }
    v32 := []int{32767, 810}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v32
    return (p1 * 0)
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]bool)
    var v3 int8 = -6
    v4 := make(map[string]string)
    fmt.Println(v1, v3)
    for fi33 := 0; fi33 < 14; fi33++ {
        switch (("\n\t" + "\n\t") + "a") {
        case "hello":
            fmt.Println(v1, fi33, v3)
        }
        v3 += v3
        _ = fi33
    }
    var v34 int64 = -1
    v3 = (^(v3))
    v35 := []int{2147483647, 42, 42, -9223372036854775808}
    switch 127 {
    case -778:
        for fi36 := 0; fi36 < 13; fi36++ {
            delete(v4, "foo bar")
            _ = fi36
        }
    case 0:
        var v37 bool = true
        _ = v37
    case 4294967295:
        for fi38 := 0; fi38 < 1; fi38++ {
            v1 -= (^(v1))
            var v39 int64 = (v34 - (v34 >> ((31) & 63)))
            _ = fi38
            _ = v39
        }
    }
    v3 = v3
    v40 := GenS2{F0: 632}
    v34 += v34
    v34 = genFn0()
    if true {
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 uint64 = 18446744073709551615
            if true {
                v43 := GenS4{F0: -1, F1: v42, F2: true}
                v35 = append(v35, fi41)
                _ = v43
            }
            _ = fi41
            _ = v42
        }
        var v44 uint = 0
        _ = v44
    }
    v34 -= (v34 ^ (v34 >> ((4294967295) & 63)))
    if (true && false) {
        v45 := GenS1{F0: (v34 + (v34 ^ -1)), F1: 529}
        fmt.Println(v34, v34)
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
}
