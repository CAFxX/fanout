package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 string, p1 string, p2 uint64) int {
    p1 = ""
    switch "" {
    case "\n\t":
        var v4 uint64 = (p2 % (((p2 % ((0) | 1))) | 1))
        _ = v4
    default:
        if ((true || (-1 > -442)) || ((0 > 2) || false)) {
            p1 = "a"
            p0 = "\n\t"
        }
    }
    if (2147483647 > 2147483647) {
        var v5 int32 = 804
        v6 := []int{0, 744}
        _ = v5
        _ = v6
    } else {
        var v7 int16 = 42
        if (false && true) {
            v7 = (v7 / ((42) | 1))
            v8 := []int{}
            if len(v8) > 0 {
                v8[0] = -604
            }
            _ = v8
        }
        _ = v7
    }
    fmt.Println(p2)
    switch 9223372036854775807 {
    case 936:
        v9 := GenS1{F0: 0, F1: 2147483647}
        _ = v9
    default:
        p0 = "foo bar"
    }
    _ = p0
    _ = p1
    _ = p2
    return -278
}

func main() {
    var v1 uint32 = 781
    v2 := []string{"", "foo bar", "\n\t"}
    var v3 int32 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: 127, F1: v1}
        if (3.14159 > 0.0) {
            v2 = append(v2, "foo bar")
            v12 := GenS2{F0: ((("foo bar" + "foo bar") + ("\n\t" + "")) + "")}
            var v13 uint64 = 941
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 bool = false
    for fi15 := 0; fi15 < 3; fi15++ {
        if ((!v14) || true) {
            for fi16 := 0; fi16 < 17; fi16++ {
                fmt.Println(v3, fi16)
                fmt.Println(v3, v1)
                _ = fi16
            }
            var v17 uint64 = 297
            fmt.Println(v1, v17, v17)
            _ = v17
        }
        var v18 uint = 4294967295
        _ = fi15
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    fmt.Println(v14)
    var v19 int64 = 712
    var v20 string = (("hello" + ("hello" + "\n\t")) + "foo bar")
    v21 := GenS3{F0: 4294967295}
    if (v14 || v14) {
        fmt.Println(v3, v1)
    }
    var v22 uint64 = 1
    var v23 int16 = 1000
    v24 := GenS1{F0: 238, F1: v1}
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(v23, v1)
            switch (v20 + "foo bar") {
            case "":
                if false {
                    if v14 {
                        v2 = append(v2, (v20 + v20))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(fi26, v19, v14)
                    } else {
                        v27 := GenS2{F0: "\n\t"}
                        fmt.Println(v19, fi25, v3)
                        _ = v27
                    }
                }
            case "hello":
                var v28 int8 = 26
                _ = v28
            case "\n\t":
                var v29 uint = 1
                _ = v29
            default:
                var v30 int8 = 1
                v19 = (v19 << ((4294967295) & 63))
                fmt.Println(v1)
                _ = len(v2)
                v1 += v1
                var v31 float32 = ((1.5 * float32(((-2.25 + 1e10) * (0.1 / 1.0)))) - 0.1)
                _ = v30
                _ = v31
            }
            _ = fi26
        }
        if false {
            v19 += (v19 << ((42) & 63))
            for fi32 := 0; fi32 < 6; fi32++ {
                _ = len(v2)
                _ = fi32
            }
        } else {
            if (2 <= 0) {
                var v33 int16 = (v23 & (v23 / ((v23) | 1)))
                _ = v33
            }
            switch v20 {
            case "\n\t":
                var v34 int32 = (v3 * v3)
                _ = v34
            }
            var v35 int8 = 0
            _ = v35
        }
        for fi36 := 0; fi36 < 8; fi36++ {
            if (true || (!(!(v14 || ((1 < 2) || (true && false)))))) {
                v37 := GenS4{F0: 0, F1: 2, F2: 1}
                _ = v37
            } else {
                fmt.Println(v3, fi36)
                switch v20 {
                case "\n\t":
                    var v38 uint8 = 2
                    _ = v38
                case "":
                    var v39 float32 = (float32(0.1) - 0.1)
                    _ = v39
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            v40 := GenS5{F0: 0, F1: -2147483648, F2: v22}
            var v41 int16 = (v23 / (((v23 ^ v23)) | 1))
            var v42 uint8 = uint8(v22)
            _ = fi36
            _ = v40
            _ = v41
            _ = v42
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
