package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int; F2 int8 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 float32 = -2.25
    var v4 float64 = 0.0
    var v5 uint32 = 2147483647
    var v6 uint16 = 32767
    fmt.Println(v3, v4, v3)
    var v7 string = "a"
    var v8 int8 = 1
    if (!true) {
        v9 := v1[-189]
        _ = v9
    }
    if (v7 >= (v7 + (v7 + (v7 + v7)))) {
        v6 = uint16(v5)
        switch 207 {
        case -9223372036854775808:
            for fi10 := 0; fi10 < 4; fi10++ {
                v3 += v3
                _ = fi10
            }
        case 127:
            v11 := GenS1{F0: -745}
            _ = v11
        case 2147483647:
            v12 := []int{}
            _ = v12
        }
    }
    var v13 bool = true
    v14 := GenS1{F0: 9223372036854775807}
    v5 -= v5
    if true {
        v15 := GenS1{F0: -1}
        v3 -= -2.25
        _ = v15
    }
    var v16 float64 = 0.1
    delete(v2, v7)
    var v17 int8 = v8
    v18 := GenS2{F0: -980, F1: 127, F2: -1}
    for fi19 := 0; fi19 < 11; fi19++ {
        if ((^(v6)) <= v6) {
            var v20 int64 = -195
            fmt.Println(v20)
            if v13 {
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 float64 = v4
                    var v23 uint32 = v5
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v24 := []int{32768}
                var v25 bool = ((32768 <= (fi19 << ((18446744073709551615) & 63))) || ((1 != 91) || (!((v13 && (!v13)) || true))))
                _ = v24
                _ = v25
            }
            _ = v20
        } else {
            switch (v7 + "hello") {
            case "a":
                v26 := []int{}
                _ = v26
            case "foo bar":
                var v27 float32 = 3.14159
                _ = v27
            case "\n\t":
                v4 = 1.5
            default:
                if (v13 || false) {
                    if true {
                        fmt.Println(v13, v13, v16)
                        for fi28 := 0; fi28 < 2; fi28++ {
                            fmt.Println(v8, v7, fi19)
                            _ = fi28
                        }
                    }
                    var v29 string = "hello"
                    _ = v29
                } else {
                    if (int64(v5) <= 42) {
                        if ((v5 < 2147483647) && ((v5 & v5) <= (v5 + (v5 >> ((29) & 63))))) {
                            v30 := GenS1{F0: (fi19 * (fi19 << ((0) & 63)))}
                            if true {
                                var v31 uint16 = 899
                                _ = v31
                            }
                            v32 := []int{(fi19 | (fi19 & (fi19 * fi19))), 128, fi19}
                            _ = v30
                            _ = v32
                        } else {
                            var v33 int64 = -965
                            _ = v33
                        }
                    } else {
                        switch (fi19 + 32767) {
                        case 0:
                            v8 += -81
                        case 191:
                            v16 -= v16
                        case 933:
                            v34 := GenS3{F0: v7}
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 4; fi35++ {
                            fmt.Println(v7)
                            _ = fi35
                        }
                    }
                    var v36 int64 = 1
                    _ = v36
                }
            }
            v37 := v1[-9223372036854775808]
            v3 = (v3 + (v3 * ((v3 * 3.14159) * (v3 / v3))))
            _ = v37
        }
        _ = fi19
    }
    switch 32767 {
    case -345:
        v38 := GenS2{F0: 42, F1: -882, F2: 0}
        _ = v38
    default:
        v39 := GenS3{F0: (v7 + v7)}
        _ = v39
    }
    v7 = v7
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
