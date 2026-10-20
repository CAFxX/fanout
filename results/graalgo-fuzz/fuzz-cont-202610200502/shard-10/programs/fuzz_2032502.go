package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
func genFn0() uint16 {
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 1
        if true {
            var v10 int8 = 42
            v9 = v9
            v11 := GenS1{F0: -1}
            _ = v10
            _ = v11
        } else {
            fmt.Println(v7, v7)
            if false {
                v12 := []int{(fi8 & (fi8 | (fi8 % ((fi8) | 1)))), fi8, (fi8 ^ -1)}
                _ = v12
            }
        }
        var v13 int8 = 127
        _ = fi8
        _ = v13
        _ = v9
    }
    switch 32768 {
    case 26:
        if true {
            v7 = v7
        }
    case -2147483648:
        fmt.Println(v7)
    default:
        for fi14 := 0; fi14 < 2; fi14++ {
            if (-32768 < 32767) {
                var v15 float32 = ((1.5 * 3.14159) - 0.1)
                _ = v15
            }
            v16 := GenS1{F0: (fi14 | (fi14 % (((^(fi14))) | 1)))}
            _ = fi14
            _ = v16
        }
    }
    _ = v7
    return (v7 * 571)
}

func main() {
    var v1 float64 = 1e10
    var v2 string = ""
    var v3 int64 = -894
    var v4 bool = false
    v5 := make(map[string]bool)
    var v6 float32 = 3.14159
    if (437 == 1) {
        var v17 bool = (v3 < v3)
        _ = v17
    }
    v2 = v2
    if ((v4 || (true || (32767 <= 42))) || true) {
        var v18 int8 = 127
        v3 = v3
        fmt.Println(v6)
        _ = v18
    }
    var v19 uint = 65535
    v19 = 983
    var v20 uint8 = uint8(v19)
    if v4 {
        var v21 uint64 = 1
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 uint8 = (v20 & 255)
            v24 := GenS1{F0: (^(fi22))}
            _ = fi22
            _ = v23
            _ = v24
        }
        switch v2 {
        case "\n\t":
            v1 = float64(1.5)
        case "a":
            switch ("" + "foo bar") {
            case "":
                switch (v2 + v2) {
                case "\n\t":
                    var v25 uint8 = (v20 - v20)
                    _ = v25
                case "":
                    for fi26 := 0; fi26 < 18; fi26++ {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            switch v2 {
                            case "a":
                                fmt.Println(v4, v6)
                            }
                            _ = fi27
                        }
                        for fi28 := 0; fi28 < 5; fi28++ {
                            v21 -= 1
                            _ = fi28
                        }
                        _ = fi26
                    }
                case "hello":
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS2{F0: (((v2 + v2) + v2) + (v2 + "")), F1: (fi29 >> ((v19) & 63)), F2: 444}
                        v31 := []int{fi29, fi29, -9223372036854775808}
                        var v32 uint8 = v20
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                default:
                    fmt.Println(v21)
                }
            }
        default:
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 uint32 = uint32(v19)
                var v35 uint8 = v20
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = v21
    } else {
        switch v2 {
        case "\n\t":
            v36 := GenS3{F0: 18446744073709551615}
            _ = v36
        case "foo bar":
            if v4 {
                v37 := GenS4{F0: 42}
                var v38 int = int(v20)
                v38 = v38
                _ = v37
                _ = v38
            }
        case "hello":
            if true {
                v4 = (v4 && v4)
                v39 := []int{}
                var v40 uint16 = 42
                _ = v39
                _ = v40
            }
        }
    }
    var v41 uint32 = 178
    fmt.Println(v6)
    var v42 uint8 = (v20 % (((v20 ^ 128)) | 1))
    for fi43 := 0; fi43 < 15; fi43++ {
        var v44 string = (("\n\t" + (v2 + ("" + ("\n\t" + (v2 + (v2 + (v2 + (("foo bar" + "hello") + v2)))))))) + "")
        _ = len(v5)
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
