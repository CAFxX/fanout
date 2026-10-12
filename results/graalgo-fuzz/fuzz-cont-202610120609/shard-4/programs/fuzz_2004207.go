package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int, p1 string, p2 string) int8 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{9223372036854775807, fi4, 9223372036854775807, -1}
        p1 = ((p1 + ("a" + p1)) + p1)
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint32 = 65535
            fmt.Println(p0, p1)
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        switch (p0 / ((32767) | 1)) {
        case -9223372036854775808:
            fmt.Println(p2, fi8, p2)
        case 32768:
            switch p0 {
            case -2147483648:
                v9 := GenS1{F0: 32767, F1: -928}
                _ = v9
            case 32767:
                for fi10 := 0; fi10 < 13; fi10++ {
                    p0 = fi8
                    var v11 int64 = -1
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        p2 = "hello"
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]uint32)
    var v3 float32 = 3.14159
    fmt.Println(v3)
    if (("hello" != ("\n\t" + "hello")) && (0 <= 738)) {
        if ((!(true || true)) || false) {
            var v13 int64 = -9223372036854775808
            v14 := GenS1{F0: 65535, F1: -796}
            var v15 int = 0
            _ = v13
            _ = v14
            _ = v15
        } else {
            var v16 uint32 = v1
            var v17 float64 = 0.0
            _ = v16
            _ = v17
        }
        var v18 float64 = (1.5 / (3.14159 + (0.0 + 1.5)))
        _ = v18
    }
    var v19 bool = ((v1 ^ v1) < 65535)
    for fi20 := 0; fi20 < 14; fi20++ {
        v1 = (v1 >> ((1) & 63))
        switch (fi20 & (^(fi20))) {
        case -110:
            if v19 {
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 float32 = float32(0.0)
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    v24 := GenS2{F0: (((("" + (("foo bar" + "a") + "\n\t")) + "hello") + "\n\t") + ("\n\t" + ("hello" + ""))), F1: v1}
                    v1 = (v1 << ((42) & 63))
                    delete(v2, -393)
                    var v25 uint32 = v1
                    v26 := GenS2{F0: "hello", F1: v1}
                    var v27 uint16 = 0
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = len(v2)
            } else {
                v28 := []int{127, (fi20 | (fi20 % (((fi20 >> ((2) & 63))) | 1))), -875}
                if v19 {
                    fmt.Println(fi20, v19)
                    delete(v2, 42)
                    v29 := GenS1{F0: 237, F1: 32767}
                    _ = v29
                } else {
                    var v30 uint64 = 18446744073709551615
                    for fi31 := 0; fi31 < 10; fi31++ {
                        if len(v28) > 0 {
                            v28[0] = (fi31 % ((-713) | 1))
                        }
                        if v19 {
                            var v32 int16 = -526
                            _ = v32
                        }
                        if (-526 == 4294967295) {
                            fmt.Println(fi20)
                        }
                        _ = fi31
                    }
                    _ = v30
                }
                _ = v28
            }
        }
        fmt.Println(v19, v19, v19)
        _ = fi20
    }
    v33 := GenS2{F0: "hello", F1: (v1 | v1)}
    switch "hello" {
    case "\n\t":
        v34 := GenS2{F0: "foo bar", F1: (v1 - 631)}
        _ = v34
    case "a":
        v3 -= v3
    case "foo bar":
        switch "a" {
        case "":
            v35 := []int{-517, 9223372036854775807, -9223372036854775808}
            _ = v35
        case "a":
            v3 += v3
        }
    }
    v36 := GenS3{F0: 127, F1: 2}
    var v37 string = "\n\t"
    v2[-9223372036854775808] = v1
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
}
