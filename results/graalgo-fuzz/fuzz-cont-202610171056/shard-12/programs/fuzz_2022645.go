package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 string, p1 int32) string {
    v7 := []int{-666, 128}
    p0 = p0
    var v8 uint = 346
    v9 := GenS1{F0: 216, F1: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 bool = true
    var v2 int16 = -343
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := []int{1}
    var v6 float64 = 1e10
    var v10 string = ""
    var v11 int64 = -9223372036854775808
    switch (((("hello" + v10) + v10) + ((v10 + v10) + "\n\t")) + "") {
    case "a":
        var v12 float32 = 3.14159
        _ = v12
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        var v13 int64 = (v11 % (((v11 + v11)) | 1))
        _ = v13
    }
    var v14 bool = ((v10 < ("hello" + v10)) && v1)
    switch "" {
    case "a":
        switch -919 {
        case 4294967295:
            if v1 {
                v14 = true
                for fi15 := 0; fi15 < 9; fi15++ {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v17 := GenS2{F0: 42}
                        if len(v5) > 0 {
                            v5[0] = (fi16 << ((0) & 63))
                        }
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 12; fi18++ {
                        _ = len(v5)
                        v11 = 9223372036854775807
                        _ = fi18
                    }
                    var v19 bool = false
                    _ = fi15
                    _ = v19
                }
                var v20 uint8 = 128
                _ = v20
            } else {
                var v21 string = (v10 + v10)
                if (!true) {
                    fmt.Println(v21, v10, v1)
                } else {
                    v3[1] = 42
                }
                _ = v21
            }
        case -864:
            if false {
                fmt.Println(v10, v6)
                var v22 uint64 = 164
                _ = v22
            }
        case 0:
            _ = len(v5)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    default:
        switch -9223372036854775808 {
        case -874:
            fmt.Println(v11, v1)
        default:
            var v23 float32 = (float32(v6) * ((0.1 + 1.5) - 0.0))
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        v10 = ((v10 + v10) + v10)
        v14 = (0 <= 171)
        var v25 uint = 0
        _ = fi24
        _ = v25
    }
    v11 += 42
    v26 := GenS1{F0: 42, F1: 1}
    switch (((("" + "hello") + (v10 + v10)) + v10) + "foo bar") {
    case "a":
        for fi27 := 0; fi27 < 10; fi27++ {
            if (int(v2) <= fi27) {
                var v28 bool = ((42 == (-(v2))) && true)
                _ = v28
            }
            _ = fi27
        }
    case "hello":
        for fi29 := 0; fi29 < 17; fi29++ {
            v11 += v11
            for fi30 := 0; fi30 < 10; fi30++ {
                if v14 {
                    var v31 float32 = -2.25
                    _ = v31
                }
                if (v11 < (-(v11))) {
                    switch (v10 + v10) {
                    case "foo bar":
                        v32 := GenS3{F0: 4294967295}
                        _ = v32
                    case "hello":
                        v33 := GenS3{F0: 213}
                        _ = v33
                    case "":
                        fmt.Println(v11)
                    }
                }
                v6 -= (1.5 + ((float64((((0.0 + 0.1) / (float32(v6) / float32(0.1))) / 1e10)) / 1.0) * 0.1))
                _ = fi30
            }
            _ = fi29
        }
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    switch v10 {
    case "a":
        var v34 float32 = 1e10
        _ = v34
    case "":
        var v35 uint64 = 18446744073709551615
        _ = v35
    }
    if (!((!(32767 <= 126)) && (true || (42 != 2147483647)))) {
        if false {
            for fi36 := 0; fi36 < 20; fi36++ {
                var v37 int64 = v11
                v10 = ((v10 + v10) + v10)
                _ = fi36
                _ = v37
            }
            v14 = (0 > 6)
            if (v1 && v1) {
                v5 = append(v5, 9223372036854775807)
                v38 := GenS3{F0: 2147483648}
                _ = v38
            }
        }
    } else {
        v5 = append(v5, -306)
    }
    if v14 {
        switch (v10 + v10) {
        case "a":
            v39 := GenS1{F0: 1, F1: 42}
            _ = v39
        }
        _ = len(v3)
    } else {
        var v40 bool = false
        v6 += v6
        var v41 uint64 = 0
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
