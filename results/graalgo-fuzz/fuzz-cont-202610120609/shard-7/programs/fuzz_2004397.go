package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 int32, p1 bool) int {
    p1 = p1
    v6 := []int{32768}
    _ = p0
    _ = p1
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 string, p2 uint32) uint16 {
    var v7 int = -613
    v8 := GenS1{F0: p2, F1: 662, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 834
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 4294967295
    var v3 string = ("foo bar" + "a")
    var v4 int64 = -297
    var v5 uint8 = 255
    v9 := []int{v1}
    v1 = (v1 >> ((42) & 63))
    v10 := GenS1{F0: 42, F1: 18446744073709551615, F2: 0}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS2{F0: 42}
        switch (^(v1)) {
        case 127:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        case 32767:
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(v5)
                if false {
                    var v14 int32 = 0
                    v2 -= v2
                    fmt.Println(v1)
                    _ = v14
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi13
            }
        case -9223372036854775808:
            switch (^(fi11)) {
            case 127:
                v9 = append(v9, 82)
            }
        }
        v15 := GenS2{F0: 4294967295}
        _ = fi11
        _ = v12
        _ = v15
    }
    fmt.Println(v1, v5)
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v1, v2)
        _ = fi16
    }
    v17 := GenS1{F0: 222, F1: (v2 - v2), F2: 127}
    var v18 string = ((v3 + ("a" + v3)) + ("a" + v3))
    v1 -= (v1 % ((-2147483648) | 1))
    if (true || (!(false && true))) {
        v19 := GenS1{F0: 0, F1: 797, F2: int8(v1)}
        var v20 uint16 = 65535
        _ = v19
        _ = v20
    } else {
        if true {
            v9 = append(v9, v1)
        }
    }
    _ = len(v9)
    v21 := GenS3{F0: v2, F1: 0}
    v4 = v4
    switch ("\n\t" + "foo bar") {
    case "a":
        if ((!true) && (true || ((18446744073709551615 != (v2 >> ((65535) & 63))) || (((1e10 + 1.5) - float32(float64(-2.25))) != 1.5)))) {
            v22 := GenS1{F0: 1, F1: v2, F2: 42}
            for fi23 := 0; fi23 < 15; fi23++ {
                fmt.Println(v3, v5, v3)
                if true {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v24 int16 = 42
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v24
                }
                var v25 int64 = v4
                _ = fi23
                _ = v25
            }
            _ = v22
        } else {
            if (1.5 > 1.5) {
                v5 = 0
                var v26 uint32 = 118
                v5 -= (^(v5))
                _ = v26
            }
            fmt.Println(v5, v18)
        }
    default:
        switch ((("\n\t" + "hello") + (v3 + "foo bar")) + v3) {
        case "\n\t":
            v27 := GenS4{F0: genFn0(-2147483648, (507 > 4294967295))}
            _ = v27
        }
    }
    v4 = v4
    v1 = v1
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int16 = 931
        if (v2 > v2) {
            v30 := GenS4{F0: fi28}
            var v31 string = v3
            for fi32 := 0; fi32 < 1; fi32++ {
                v9 = append(v9, (^(v1)))
                if len(v9) > 0 {
                    v9[0] = 1
                }
                v1 -= fi32
                _ = fi32
            }
            _ = v30
            _ = v31
        } else {
            v33 := GenS5{F0: v18, F1: genFn1(1, "", 94)}
            var v34 string = (v18 + "\n\t")
            _ = v33
            _ = v34
        }
        _ = fi28
        _ = v29
    }
    var v35 float32 = float32(((-2.25 / (0.1 - -2.25)) * 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
