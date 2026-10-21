package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32 }
func genFn0() string {
    v4 := []int{749, 128}
    var v5 uint16 = 0
    _ = v4
    _ = v5
    return ""
}
func genFn1() bool {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int64 = 0
        var v8 int8 = 42
        switch ((("hello" + "\n\t") + ("a" + "foo bar")) + ((("\n\t" + ("hello" + "hello")) + "\n\t") + "foo bar")) {
        case "":
            var v9 int32 = 1
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 32768 {
    case 127:
        var v10 int32 = -516
        _ = v10
    default:
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint = 196
            v13 := []int{fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 uint64 = 651
    v15 := GenS1{F0: -32768, F1: -2147483648}
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 int = (fi16 * (fi16 & fi16))
        if true {
            for fi18 := 0; fi18 < 4; fi18++ {
                if true {
                    v19 := GenS2{F0: 18446744073709551615, F1: 42, F2: 65535}
                    v14 = 2
                    var v20 float64 = ((((3.14159 * 0.0) + 0.1) - (float64(1e10) + (3.14159 + (3.14159 + -2.25)))) + float64(1e10))
                    _ = v19
                    _ = v20
                }
                v21 := []int{(-(fi16))}
                _ = fi18
                _ = v21
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 uint32 = 1
                if true {
                    switch "hello" {
                    case "a":
                        var v24 uint32 = 221
                        _ = v24
                    case "\n\t":
                        var v25 uint64 = v14
                        _ = v25
                    case "":
                        v26 := GenS3{F0: (fi16 >> ((812) & 63)), F1: 991}
                        _ = v26
                    }
                    fmt.Println(v14, v14)
                } else {
                    v27 := []int{fi16, -9223372036854775808, fi22}
                    _ = v27
                }
                if (1 <= (v23 | v23)) {
                    v14 += uint64(v23)
                    var v28 float32 = float32((0.1 + -2.25))
                    _ = v28
                } else {
                    var v29 int8 = -1
                    v30 := GenS2{F0: 0, F1: 0, F2: 65535}
                    v31 := GenS4{F0: 32767}
                    var v32 int = fi16
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v14 -= (v14 + v14)
                if false {
                    v17 = (v17 * fi16)
                }
                _ = fi22
                _ = v23
            }
            var v33 uint16 = 65535
            _ = v33
        }
        v34 := []int{fi16, 998}
        _ = fi16
        _ = v17
        _ = v34
    }
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 uint16 = 139
    v2 := make(map[string]int)
    var v3 int16 = 32767
    _ = len(v2)
    v35 := GenS4{F0: v3}
    v36 := GenS1{F0: int16(v1), F1: -141}
    var v37 float64 = (0.0 * (0.0 + 0.0))
    switch "a" {
    case "\n\t":
        v37 = v37
    case "":
        switch (("hello" + ("\n\t" + "")) + "foo bar") {
        case "hello":
            for fi38 := 0; fi38 < 2; fi38++ {
                v1 = v1
                _ = fi38
            }
        }
    case "foo bar":
        var v39 int64 = 0
        _ = v39
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        v3 = v3
        for fi41 := 0; fi41 < 8; fi41++ {
            v37 = v37
            v42 := GenS5{F0: 2147483647}
            fmt.Println(v37)
            _ = fi41
            _ = v42
        }
        switch -9223372036854775808 {
        case 32768:
            if (42 != -634) {
                v43 := []int{-1}
                v37 -= 1e10
                _ = v43
            }
        case -565:
            var v44 float64 = float64(1.5)
            _ = v44
        case 4294967295:
            v1 -= (v1 ^ v1)
        }
        _ = fi40
    }
    if (2 != 64) {
        var v45 uint32 = 514
        fmt.Println(v3)
        _ = v45
    } else {
        var v46 int32 = 42
        _ = v46
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 int = (fi47 & 2147483647)
        _ = fi47
        _ = v48
    }
    if true {
        if (v37 < v37) {
            v3 = 32767
            if true {
                v37 -= v37
                var v49 uint8 = 42
                v50 := GenS4{F0: (v3 / ((v3) | 1))}
                _ = v49
                _ = v50
            }
        }
        var v51 uint = 4294967295
        _ = v51
    } else {
        v37 = v37
    }
    v1 = (v1 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
