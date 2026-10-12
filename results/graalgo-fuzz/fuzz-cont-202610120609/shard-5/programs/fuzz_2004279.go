package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint16, p1 uint, p2 uint8) int {
    var v6 uint8 = p2
    v7 := []int{2147483647, -2147483648, -937, -406}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(p2, p1)
        _ = fi8
    }
    p1 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := []uint32{2147483647, 243}
    v2 := []uint32{4294967295, 0}
    var v3 int16 = 833
    var v4 int32 = -1
    v5 := make(map[string]int)
    v4 += (v4 + v4)
    v9 := GenS1{F0: 42}
    var v10 uint16 = 32767
    v1 = append(v1, 868)
    var v11 uint8 = 2
    v10 = v10
    fmt.Println(v10)
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(fi12)
        switch ("a" + "") {
        case "foo bar":
            delete(v5, "foo bar")
        case "":
            var v13 uint32 = 570
            _ = v13
        case "hello":
            v10 += 2
        }
        _ = fi12
    }
    v14 := v5[""]
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v10, v10)
        if (42 == 0) {
            if (!(true && ((false || (!(fi15 <= (fi15 / (((fi15 % ((v14) | 1))) | 1))))) && ("a" <= (((("\n\t" + (("hello" + "hello") + "foo bar")) + ("foo bar" + "foo bar")) + "") + ("\n\t" + "")))))) {
                var v16 uint = 0
                v17 := v5[""]
                _ = v16
                _ = v17
            }
            v14 = v14
        } else {
            switch "\n\t" {
            case "foo bar":
                var v18 int64 = 42
                _ = v18
            case "a":
                var v19 float32 = 0.0
                _ = v19
            case "hello":
                v4 = (v4 & v4)
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                if false {
                    var v21 string = "a"
                    _ = v21
                }
                var v22 int32 = (v4 - v4)
                _ = fi20
                _ = v22
            }
        }
        _ = fi15
    }
    fmt.Println(v11)
    var v23 float64 = 0.1
    for fi24 := 0; fi24 < 2; fi24++ {
        v10 -= v10
        v14 = -966
        _ = fi24
    }
    v3 -= (v3 ^ 600)
    for fi25 := 0; fi25 < 4; fi25++ {
        if (!(1 == uint(v10))) {
            var v26 uint8 = v11
            _ = v26
        }
        switch (-(fi25)) {
        case 2147483647:
            switch (v14 - v14) {
            case 32768:
                v3 += (v3 * int16(v10))
            default:
                var v27 int32 = (v4 * v4)
                _ = v27
            }
        case 32768:
            v28 := GenS1{F0: 65535}
            _ = v28
        default:
            v23 = v23
        }
        switch "foo bar" {
        case "a":
            v5["\n\t"] = v14
        case "\n\t":
            var v29 uint32 = 42
            _ = v29
        case "hello":
            switch "\n\t" {
            case "hello":
                v14 += (v14 * fi25)
            case "foo bar":
                fmt.Println(fi25, v11)
            case "\n\t":
                if true {
                    if ("" <= "hello") {
                        v4 = (v4 % (((^(v4))) | 1))
                        if false {
                            v3 -= -965
                            v3 = v3
                        } else {
                            if false {
                                v10 -= (^(v10))
                                var v30 uint32 = 782
                                v31 := GenS2{F0: 263, F1: int32(v30)}
                                _ = v30
                                _ = v31
                            }
                            var v32 int = int(v3)
                            _ = v32
                        }
                        v5[""] = 522
                    } else {
                        v33 := GenS3{F0: (("foo bar" + ("hello" + ("" + "\n\t"))) + (("" + ("" + "")) + "a")), F1: (v4 % ((v4) | 1))}
                        v34 := GenS4{F0: 1, F1: (v11 * (v11 - v11))}
                        _ = v33
                        _ = v34
                    }
                } else {
                    fmt.Println(v10)
                    v14 += (v14 >> ((7) & 63))
                }
            default:
                switch "" {
                case "foo bar":
                    var v35 int = 263
                    _ = v35
                }
            }
        }
        _ = fi25
    }
    v36 := GenS1{F0: 65535}
    var v37 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
