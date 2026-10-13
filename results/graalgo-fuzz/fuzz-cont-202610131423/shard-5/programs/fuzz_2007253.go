package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int16; F2 string }
func genFn0(p0 int32, p1 uint8) string {
    for fi4 := 0; fi4 < 16; fi4++ {
        if (true && true) {
            if false {
                for fi5 := 0; fi5 < 10; fi5++ {
                    v6 := []int{200}
                    _ = len(v6)
                    v6 = append(v6, 127)
                    _ = fi5
                    _ = v6
                }
                var v7 uint16 = 42
                _ = v7
            }
            var v8 bool = (164 <= p1)
            _ = v8
        } else {
            var v9 int64 = 1
            for fi10 := 0; fi10 < 12; fi10++ {
                fmt.Println(v9, p0, fi4)
                _ = fi10
            }
            if true {
                p1 = 1
            } else {
                fmt.Println(fi4, p0)
            }
            _ = v9
        }
        var v11 float64 = ((1e10 + ((float64(1e10) * -2.25) - 3.14159)) - (1.5 + 1.5))
        _ = fi4
        _ = v11
    }
    if true {
        v12 := []int{2147483647}
        _ = v12
    }
    var v13 int = -746
    p1 -= 42
    _ = p0
    _ = p1
    _ = v13
    return "hello"
}
func genFn1(p0 int32, p1 int32, p2 bool) int16 {
    if p2 {
        v14 := []int{-2147483648, -686}
        _ = v14
    } else {
        switch "foo bar" {
        case "hello":
            fmt.Println(p1)
        default:
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 int8 = 83
                    fmt.Println(p0, p2, fi15)
                    p1 = (p1 - 1)
                    _ = fi16
                    _ = v17
                }
                var v18 uint16 = 1
                _ = fi15
                _ = v18
            }
        }
    }
    var v19 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 1
    var v3 int = -799
    v20 := GenS1{F0: (((v3 - (v3 + (^(v3)))) <= 127) || false), F1: "hello"}
    v3 = (v3 - (v3 * (v3 - -9223372036854775808)))
    v21 := []int{(v3 >> ((0) & 63)), v3}
    var v22 uint = 1
    var v23 string = ("hello" + "foo bar")
    var v24 int64 = 238
    if true {
        switch int(v22) {
        case 9223372036854775807:
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS2{F0: v2}
                var v27 string = ("" + v23)
                _ = fi25
                _ = v26
                _ = v27
            }
        case 42:
            for fi28 := 0; fi28 < 11; fi28++ {
                switch "" {
                case "a":
                    v3 = (-(fi28))
                }
                v22 += (v22 << (((v22 | v22)) & 63))
                switch ((v23 + v23) + v23) {
                case "a":
                    v22 -= (v22 / (((v22 + 1)) | 1))
                case "":
                    v29 := GenS3{F0: 1, F1: 42, F2: ((v23 + v23) + v23)}
                    _ = v29
                case "hello":
                    delete(v1, 1)
                default:
                    for fi30 := 0; fi30 < 10; fi30++ {
                        if len(v21) > 0 {
                            v21[0] = (fi30 & 641)
                        }
                        if ((568 != 32767) || true) {
                            if true {
                                v3 += 2147483647
                                if (4294967295 < 2147483647) {
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                } else {
                                    var v31 int = fi30
                                    var v32 int32 = 0
                                    _ = v31
                                    _ = v32
                                }
                                v33 := GenS3{F0: 2147483647, F1: genFn1(1, -486, ((0 < 2) && ((((1 >= -1) && (!false)) || true) && false))), F2: ("\n\t" + (((v23 + "foo bar") + v23) + "a"))}
                                _ = v33
                            }
                            v34 := GenS3{F0: 0, F1: -32768, F2: (v23 + (v23 + v23))}
                            _ = v34
                        } else {
                            var v35 int8 = 100
                            v3 += (v3 / ((-1) | 1))
                            _ = v35
                        }
                        v36 := GenS1{F0: ((v22 % ((v22) | 1)) > 255), F1: ("\n\t" + "a")}
                        _ = fi30
                        _ = v36
                    }
                }
                _ = fi28
            }
        case -1:
            switch 127 {
            case 9223372036854775807:
                if (!true) {
                    v24 = -1
                } else {
                    var v37 uint64 = 1
                    _ = v37
                }
            case 2147483647:
                fmt.Println(v24)
            }
        default:
            v38 := GenS2{F0: (v2 | v2)}
            _ = v38
        }
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
