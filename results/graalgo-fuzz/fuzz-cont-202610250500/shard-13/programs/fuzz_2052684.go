package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 int8) int16 {
    fmt.Println(p1, p1)
    if p0 {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 float32 = -2.25
            for fi8 := 0; fi8 < 3; fi8++ {
                p1 += p1
                switch 127 {
                case 127:
                    if p0 {
                        var v9 string = "hello"
                        var v10 float32 = v7
                        _ = v10
                        _ = v9
                    } else {
                        v11 := GenS1{F0: 42, F1: 1000, F2: -15}
                        var v12 bool = (!(p0 || p0))
                        p0 = (119 == fi8)
                        _ = v11
                        _ = v12
                    }
                case 1:
                    var v13 int = (fi6 - 4294967295)
                    _ = v13
                case -9223372036854775808:
                    v14 := []int{fi6, 42, 9223372036854775807}
                    _ = v14
                }
                if ((((-2.25 - 3.14159) + (0.0 / 3.14159)) * 0.0) > float32(float64(v7))) {
                    if (!(0 >= (-(fi6)))) {
                        v15 := []int{fi6, fi8, (-(fi8))}
                        _ = len(v15)
                        fmt.Println(fi8, p0, p1)
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        v15 = append(v15, 2147483647)
                        var v16 bool = false
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 1, F1: 1, F2: -53}
                        p1 = p1
                        v18 := GenS1{F0: 774, F1: 224, F2: (-(p1))}
                        _ = v17
                        _ = v18
                    }
                } else {
                    if (p1 > -1) {
                        var v19 uint64 = 42
                        var v20 uint64 = (^(v19))
                        _ = v19
                        _ = v20
                    }
                    switch ((("a" + ("" + "a")) + "\n\t") + "hello") {
                    case "a":
                        fmt.Println(fi8, p0)
                    case "\n\t":
                        var v21 int32 = 2147483647
                        _ = v21
                    }
                    var v22 uint64 = 1
                    _ = v22
                }
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
    }
    var v23 string = "foo bar"
    switch -9223372036854775808 {
    case 127:
        var v24 int8 = (p1 & 1)
        _ = v24
    case -1:
        switch 721 {
        case 9223372036854775807:
            for fi25 := 0; fi25 < 15; fi25++ {
                fmt.Println(p0, fi25, v23)
                _ = fi25
            }
        }
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v23
    return 1000
}

func main() {
    v1 := []string{"\n\t", ("hello" + (("hello" + "foo bar") + ("\n\t" + ""))), ("" + "hello"), "\n\t"}
    var v2 int32 = 0
    var v3 float32 = 0.1
    var v4 bool = (true && ("hello" <= ("\n\t" + ("a" + "\n\t"))))
    var v5 int64 = 0
    v26 := GenS1{F0: -553, F1: 87, F2: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v4 {
        switch "a" {
        case "hello":
            var v27 int64 = (v5 << ((4294967295) & 63))
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(v5)
            var v29 float32 = 0.0
            if v4 {
                v2 = (^(v2))
                var v30 int64 = (^(v5))
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = v30
            } else {
                var v31 int8 = 127
                v32 := GenS1{F0: 42, F1: int16(v2), F2: (v31 - -46)}
                var v33 bool = v4
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi28
            _ = v29
        }
    }
    _ = len(v1)
    v34 := GenS1{F0: genFn0(((v4 || false) || false), 42), F1: genFn0((!(v4 && false)), 1), F2: 100}
    var v35 uint32 = 4294967295
    var v36 int = -317
    v37 := GenS2{F0: v2, F1: 42}
    var v38 int8 = -1
    if true {
        switch -97 {
        case 32767:
            if (v4 || true) {
                v35 -= v35
                v35 = v35
            }
        }
        if len(v1) > 0 {
            v1[0] = ""
        }
        v4 = ((v35 ^ v35) <= (v35 & 2147483648))
    }
    if v4 {
        v39 := GenS3{F0: ("" + "\n\t")}
        v40 := GenS2{F0: (v2 % (((v2 >> ((31) & 63))) | 1)), F1: 18446744073709551615}
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
