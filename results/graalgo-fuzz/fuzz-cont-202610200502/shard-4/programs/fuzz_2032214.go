package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 int }
func genFn0() int {
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5)
        for fi6 := 0; fi6 < 15; fi6++ {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                var v7 string = "foo bar"
                _ = v7
            case "a":
                var v8 uint16 = 434
                _ = v8
            case "":
                fmt.Println(fi5, fi5)
                var v9 uint8 = 74
                v10 := []int{(fi6 + (-(fi5))), fi6, fi6}
                if len(v10) > 0 {
                    v10[0] = 2147483647
                }
                var v11 int32 = 1
                _ = v10
                _ = v11
                _ = v9
            }
            switch (("\n\t" + ("\n\t" + "hello")) + ("a" + "\n\t")) {
            case "a":
                var v12 string = "hello"
                _ = v12
            case "hello":
                var v13 float64 = ((0.1 * (-2.25 + ((0.0 + (-2.25 + (float64((1.5 - 1.5)) / (1.5 + 0.1)))) - float64(1e10)))) / 1e10)
                _ = v13
            case "\n\t":
                v14 := []int{fi6}
                _ = v14
            }
            var v15 float32 = -2.25
            _ = fi6
            _ = v15
        }
        if (1 != 0) {
            if (67 != -65) {
                var v16 int64 = 430
                fmt.Println(fi5)
                _ = v16
            }
            fmt.Println(fi5)
        } else {
            var v17 int = (fi5 / (((fi5 - fi5)) | 1))
            switch "" {
            case "":
                if (42 <= 0) {
                    v18 := []int{v17, fi5}
                    _ = v18
                } else {
                    switch ("a" + "a") {
                    case "foo bar":
                        v19 := GenS1{F0: -2147483648, F1: -9223372036854775808, F2: false}
                        _ = v19
                    default:
                        if ((false && true) && true) {
                            var v20 float32 = -2.25
                            _ = v20
                        } else {
                            var v21 int64 = 1
                            v22 := GenS1{F0: 0, F1: v21, F2: false}
                            _ = v21
                            _ = v22
                        }
                    }
                    if (2147483647 < 722) {
                        v17 = (v17 / ((-156) | 1))
                        var v23 uint16 = 32767
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v17 = (v17 >> ((4294967295) & 63))
                            var v25 int16 = -32768
                            _ = fi24
                            _ = v25
                        }
                    }
                }
            case "foo bar":
                v26 := []int{(v17 | (fi5 ^ fi5))}
                _ = v26
            default:
                var v27 string = (("foo bar" + "hello") + "foo bar")
                _ = v27
            }
            _ = v17
        }
        _ = fi5
    }
    var v28 int32 = 42
    _ = v28
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -59
    var v3 bool = false
    var v4 int64 = -1
    v29 := GenS2{F0: 2, F1: 42, F2: 32767}
    if ((-(v4)) > (v4 & v4)) {
        var v30 int8 = -122
        var v31 uint = 42
        fmt.Println(v31, v3)
        _ = v30
        _ = v31
    }
    fmt.Println(v2, v4)
    var v32 bool = true
    if v3 {
        var v33 int16 = -1
        v2 += v2
        v3 = (false && true)
        _ = v33
    }
    switch "" {
    case "\n\t":
        if (!((859 <= 0) && v3)) {
            var v34 uint32 = 2147483647
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 string = "hello"
                _ = fi35
                _ = v36
            }
            for fi37 := 0; fi37 < 17; fi37++ {
                fmt.Println(v34, v3, fi37)
                _ = fi37
            }
            _ = v34
        } else {
            v4 = v4
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 uint = 255
                _ = len(v1)
                var v40 int16 = 968
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    case "hello":
        var v41 uint8 = 109
        _ = v41
    default:
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 bool = (1 != v2)
            var v44 int = fi42
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    v45 := GenS3{F0: 0}
    if ((!(42 < 146)) && true) {
        v46 := GenS2{F0: 48, F1: 26, F2: 32768}
        var v47 uint32 = 2
        _ = v46
        _ = v47
    }
    v2 += -102
    switch (("foo bar" + "\n\t") + "") {
    case "\n\t":
        _ = len(v1)
    }
    delete(v1, 2147483647)
    for fi48 := 0; fi48 < 3; fi48++ {
        v4 += v4
        _ = fi48
    }
    if (2 < 126) {
        fmt.Println(v3, v2, v4)
    } else {
        switch 32767 {
        case 128:
            switch ("hello" + "hello") {
            case "hello":
                var v49 uint16 = 65535
                _ = v49
            }
        }
    }
    v1[-2147483648] = 4294967295
    var v50 float32 = ((1.5 * ((-2.25 / float32(1.5)) / 1.5)) / 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v45
    _ = v50
}
