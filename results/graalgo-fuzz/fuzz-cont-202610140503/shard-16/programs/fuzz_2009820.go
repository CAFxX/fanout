package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64) uint16 {
    v4 := GenS1{F0: 2, F1: 154, F2: 42}
    p0 = p0
    var v5 float32 = ((1e10 + 3.14159) + float32(1e10))
    v5 = (v5 + v5)
    switch 535 {
    case 9223372036854775807:
        var v6 bool = false
        _ = v6
    case 797:
        v7 := GenS1{F0: 0, F1: 180, F2: 95}
        _ = v7
    case 2147483647:
        v8 := []int{-566, 4294967295}
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []uint32{712, 42}
    v2 := make(map[string]string)
    v3 := []uint32{2147483647, 2}
    if (false || true) {
        var v9 uint16 = 32767
        for fi10 := 0; fi10 < 18; fi10++ {
            v9 = v9
            v11 := GenS1{F0: 2, F1: (^(v9)), F2: 18446744073709551615}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            if (false || ((!false) || (false && true))) {
                var v13 uint16 = v9
                _ = v13
            }
            var v14 uint16 = v9
            _ = fi12
            _ = v14
        }
        _ = v9
    } else {
        switch "a" {
        case "":
            if (!true) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                for fi15 := 0; fi15 < 14; fi15++ {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        _ = len(v1)
                        v17 := GenS2{F0: 1}
                        var v18 int = fi16
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v19 uint32 = 42
                    _ = fi15
                    _ = v19
                }
            }
        case "\n\t":
            var v20 int = 9223372036854775807
            _ = v20
        case "hello":
            var v21 uint64 = 69
            _ = v21
        }
    }
    var v22 uint8 = 255
    var v23 uint32 = 2147483647
    var v24 bool = true
    v24 = (42 < 1)
    v24 = true
    fmt.Println(v23, v24, v22)
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(v24, v22)
        _ = fi25
    }
    v24 = (42 > 4294967295)
    var v26 int8 = 39
    var v27 int16 = -1
    var v28 uint = 65535
    if (float64(1.5) > ((-2.25 + -2.25) + ((-2.25 * 3.14159) * -2.25))) {
        var v29 int16 = v27
        if (!true) {
            var v30 uint32 = v23
            if (v29 == 42) {
                var v31 string = ("hello" + ("hello" + ("foo bar" + (("hello" + "foo bar") + "hello"))))
                _ = v31
            } else {
                if v24 {
                    v22 = 0
                    v32 := GenS1{F0: 128, F1: genFn0(-1), F2: 612}
                    _ = v32
                }
                if v24 {
                    v33 := GenS3{F0: 18446744073709551615}
                    v34 := GenS2{F0: 32767}
                    _ = v33
                    _ = v34
                }
                if (v24 && ((935 != 1) || v24)) {
                    var v35 float64 = -2.25
                    _ = v35
                }
            }
            v36 := GenS1{F0: (v22 << ((v28) & 63)), F1: genFn0(int64(v26)), F2: 42}
            _ = v30
            _ = v36
        }
        _ = v29
    } else {
        var v37 string = ("foo bar" + "a")
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 int8 = int8(v22)
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 uint = 42
                var v42 float64 = 1.5
                if (!(170 >= 0)) {
                    fmt.Println(v22, v37)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v43 uint32 = 2
                    _ = v43
                }
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi38
            _ = v39
        }
        fmt.Println(v26)
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
