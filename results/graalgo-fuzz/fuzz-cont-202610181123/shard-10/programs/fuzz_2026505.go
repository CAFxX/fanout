package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint }
func genFn0() int8 {
    var v4 float64 = 3.14159
    switch 4294967295 {
    case 128:
        var v5 bool = true
        _ = v5
    case 849:
        var v6 float32 = -2.25
        _ = v6
    case 9223372036854775807:
        v7 := []int{9223372036854775807, -850, -9223372036854775808, 127}
        _ = v7
    }
    _ = v4
    return 127
}

func main() {
    var v1 int64 = -615
    v2 := []string{("\n\t" + (("hello" + "hello") + "a"))}
    v3 := make(map[string]int)
    v8 := GenS1{F0: 978, F1: 1, F2: 185}
    var v9 int16 = 1000
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 = v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 513, F1: 127, F2: 2}
        fmt.Println(fi11, v1, v9)
        v13 := GenS2{F0: (-(fi11))}
        _ = fi11
        _ = v12
        _ = v13
    }
    v9 = (v9 ^ v9)
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS2{F0: fi15}
            fmt.Println(fi14, v1)
            var v17 float64 = (1.5 + 0.1)
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 uint8 = 42
            _ = fi18
            _ = v19
        }
        if (0.0 < (3.14159 - (0.1 - (0.1 * 0.1)))) {
            if (!true) {
                if false {
                    if ((0 > 1) && (!(false || (18446744073709551615 != 1)))) {
                        v20 := v3["a"]
                        _ = v20
                    } else {
                        _ = len(v3)
                        var v21 float64 = (1e10 + float64((((float32(0.0) + 0.0) / 1.0) * float32((3.14159 / -2.25)))))
                        _ = v21
                    }
                }
            } else {
                if true {
                    fmt.Println(v1, v1, v9)
                }
                if false {
                    if true {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v23 := GenS1{F0: 1, F1: 2, F2: 128}
                            _ = fi22
                            _ = v23
                        }
                        var v24 float32 = 0.1
                        _ = v24
                    }
                    v25 := GenS2{F0: 42}
                    _ = v25
                } else {
                    var v26 int64 = v1
                    var v27 int8 = 1
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 bool = false
                    v29 = v29
                    v30 := GenS1{F0: 0, F1: 251, F2: 127}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        } else {
            fmt.Println(v9, v1)
        }
        _ = fi14
    }
    if false {
        fmt.Println(v9, v9)
    } else {
        var v31 bool = true
        var v32 uint64 = 0
        _ = v31
        _ = v32
    }
    v9 -= -32768
    for fi33 := 0; fi33 < 16; fi33++ {
        v9 += (v9 | v9)
        var v34 int = fi33
        _ = fi33
        _ = v34
    }
    if true {
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := v3[""]
                if true {
                    fmt.Println(v9, v36)
                } else {
                    v9 += 1000
                }
                _ = fi35
                _ = v36
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (890 > 1) {
            for fi37 := 0; fi37 < 10; fi37++ {
                v1 = 1
                switch "foo bar" {
                case "\n\t":
                    v9 -= v9
                case "":
                    var v38 int16 = v9
                    _ = v38
                case "hello":
                    switch (fi37 & (fi37 % ((int(v9)) | 1))) {
                    case 2147483647:
                        v39 := GenS3{F0: 0, F1: 833}
                        _ = v39
                    default:
                        fmt.Println(v9, fi37, v9)
                    }
                }
                _ = fi37
            }
            switch ("a" + (("hello" + (("" + "") + "foo bar")) + "a")) {
            case "\n\t":
                switch "a" {
                case "\n\t":
                    var v40 uint8 = 255
                    _ = v40
                }
            }
        }
    }
    v3[("foo bar" + "hello")] = -1
    if true {
        var v41 uint = 116
        _ = v41
    } else {
        v1 += (v1 / ((-512) | 1))
        fmt.Println(v1)
        v9 -= (v9 >> ((18446744073709551615) & 63))
    }
    v42 := GenS3{F0: -89, F1: 1}
    v43 := GenS3{F0: 100, F1: 552}
    fmt.Println(v1)
    var v44 bool = false
    var v45 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v8
    _ = v9
}
