package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int32; F1 int8; F2 int32 }

func main() {
    var v1 int16 = 1000
    v2 := []string{"hello", (("foo bar" + ("hello" + "")) + "hello"), "\n\t"}
    var v3 int32 = -1
    _ = len(v2)
    _ = len(v2)
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := GenS1{F0: v1}
            v2 = append(v2, "foo bar")
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v3 = v3
    if true {
        if len(v2) > 0 {
            v2[0] = (("hello" + "\n\t") + "")
        }
        v1 += v1
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            if (-900 <= (v3 + 350)) {
                if false {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        if (42 != 172) {
                            fmt.Println(fi7, v1)
                        }
                        _ = fi8
                    }
                    for fi9 := 0; fi9 < 17; fi9++ {
                        v10 := GenS1{F0: v1}
                        switch ((("a" + "\n\t") + "") + "foo bar") {
                        case "a":
                            fmt.Println(fi7)
                        default:
                            var v11 uint64 = 1
                            _ = v11
                        }
                        _ = fi9
                        _ = v10
                    }
                }
            }
            v12 := GenS1{F0: (v1 ^ (v1 % ((v1) | 1)))}
            _ = fi7
            _ = v12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                v1 -= 0
                _ = fi14
            }
            v2 = append(v2, "\n\t")
            _ = fi13
        }
        v3 -= -974
    }
    v15 := GenS1{F0: v1}
    if (!true) {
        v16 := GenS2{F0: 2147483648, F1: 1}
        fmt.Println(v1)
        if false {
            if len(v2) > 0 {
                v2[0] = ("hello" + ("" + "a"))
            }
            var v17 int16 = (v1 * (^(v1)))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v17
        }
        _ = v16
    }
    fmt.Println(v3)
    fmt.Println(v3, v3)
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            var v19 int32 = 498
            var v20 uint64 = 18446744073709551615
            _ = v19
            _ = v20
        }
        var v21 int = fi18
        if (0 != (v1 << ((4294967295) & 63))) {
            if false {
                var v22 uint16 = 0
                var v23 string = "a"
                _ = v22
                _ = v23
            }
            var v24 uint8 = 0
            _ = v24
        } else {
            for fi25 := 0; fi25 < 17; fi25++ {
                switch (fi18 >> ((286) & 63)) {
                case -1:
                    fmt.Println(v21)
                default:
                    v21 = v21
                }
                _ = fi25
            }
            fmt.Println(v3, v3, v3)
            v3 -= 1
        }
        _ = fi18
        _ = v21
    }
    v26 := GenS1{F0: 1000}
    v27 := GenS1{F0: (v1 ^ (v1 + 32767))}
    v28 := GenS3{F0: v3, F1: 1, F2: (v3 ^ v3)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "hello":
        var v29 float32 = 1e10
        _ = v29
    case "foo bar":
        switch "a" {
        case "hello":
            switch "hello" {
            case "a":
                v3 += v3
            }
        }
    default:
        switch "hello" {
        case "a":
            for fi30 := 0; fi30 < 9; fi30++ {
                for fi31 := 0; fi31 < 8; fi31++ {
                    _ = len(v2)
                    v1 = -32768
                    _ = fi31
                }
                v3 = (^(v3))
                var v32 int64 = 42
                _ = fi30
                _ = v32
            }
        }
    }
    if true {
        v33 := GenS3{F0: 2147483647, F1: -1, F2: (v3 | -630)}
        v3 = (v3 & -1)
        _ = v33
    } else {
        var v34 uint64 = 18446744073709551615
        _ = len(v2)
        v35 := GenS3{F0: 0, F1: 1, F2: 2147483647}
        _ = v34
        _ = v35
    }
    var v36 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
}
