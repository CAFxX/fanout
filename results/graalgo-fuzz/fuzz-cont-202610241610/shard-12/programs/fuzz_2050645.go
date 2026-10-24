package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint; F2 int16 }
func genFn0() int16 {
    var v5 uint32 = 1
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 4294967295}
            switch 32767 {
            case 32768:
                v5 = 620
            case -9223372036854775808:
                fmt.Println(fi6)
                v5 += v5
            case 0:
                v5 += v5
                v9 := []int{}
                fmt.Println(fi7, fi7)
                _ = v9
            default:
                var v10 float64 = 1e10
                _ = v10
            }
            v11 := GenS1{F0: 2}
            _ = fi7
            _ = v11
            _ = v8
        }
        fmt.Println(fi6, fi6)
        v12 := GenS1{F0: 2147483648}
        _ = fi6
        _ = v12
    }
    var v13 int64 = 0
    _ = v13
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 2
    fmt.Println(v4)
    if false {
        var v14 uint = 42
        _ = v14
    }
    _ = len(v3)
    v4 += v4
    var v15 float32 = -2.25
    for fi16 := 0; fi16 < 6; fi16++ {
        v15 = 3.14159
        switch "foo bar" {
        case "a":
            if false {
                var v17 uint32 = 42
                _ = v17
            } else {
                v4 -= (v4 * 0)
                v18 := GenS2{F0: 1, F1: 644, F2: 32767}
                var v19 string = "\n\t"
                _ = v18
                _ = v19
            }
        case "foo bar":
            var v20 uint8 = v4
            _ = v20
        }
        v4 = (^(v4))
        _ = fi16
    }
    v21 := []int{}
    v15 = (v15 + v15)
    if ((v4 > (^(v4))) || ((42 == 104) && (true || (106 <= -1)))) {
        var v22 uint64 = 0
        switch int(v4) {
        case 42:
            v23 := GenS1{F0: 42}
            _ = v23
        case 9223372036854775807:
            fmt.Println(v22)
        case -9223372036854775808:
            v21 = append(v21, -2147483648)
        }
        if (!(false && true)) {
            v24 := GenS1{F0: 4294967295}
            if false {
                switch 42 {
                case 0:
                    var v25 string = ("hello" + "a")
                    _ = v25
                }
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                if true {
                    v22 += v22
                    v15 = -2.25
                    var v26 int16 = 1000
                    _ = v26
                } else {
                    switch int(v4) {
                    case 32767:
                        v22 += (^(v22))
                    case 42:
                        v27 := v3[1]
                        _ = v27
                    }
                }
            }
            _ = v24
        } else {
            var v28 int = -1
            var v29 uint = 1
            var v30 int64 = 975
            _ = v28
            _ = v29
            _ = v30
        }
        _ = v22
    }
    if false {
        if true {
            if true {
                v31 := GenS2{F0: 18446744073709551615, F1: 0, F2: 630}
                _ = v31
            }
            var v32 float64 = 1.5
            v33 := GenS3{F0: 0, F1: 4294967295, F2: -32768}
            _ = v32
            _ = v33
        }
    }
    v34 := GenS3{F0: -948, F1: 579, F2: 32767}
    v35 := v2["hello"]
    fmt.Println(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
