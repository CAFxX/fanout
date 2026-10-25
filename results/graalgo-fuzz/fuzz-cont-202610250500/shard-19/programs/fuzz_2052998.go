package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 int64 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        if (fi5 >= 987) {
            var v6 float32 = 1.5
            _ = v6
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 2}
            _ = fi7
            _ = v8
        }
        var v9 int = fi5
        _ = fi5
        _ = v9
    }
    var v10 int16 = 42
    v10 -= -864
    var v11 uint32 = 289
    v12 := []int{128}
    _ = v10
    _ = v11
    _ = v12
    return v10
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 uint = 4294967295
    v13 := GenS2{F0: 0, F1: 38, F2: 1}
    var v14 int8 = 0
    for fi15 := 0; fi15 < 13; fi15++ {
        if true {
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(fi16, fi16, v14)
                var v17 int16 = -260
                _ = fi16
                _ = v17
            }
            v18 := GenS3{F0: 93, F1: 42, F2: 42}
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := []int{9223372036854775807}
                if (820 == 42) {
                    fmt.Println(fi15)
                    var v21 int64 = 9223372036854775807
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        _ = fi15
    }
    v22 := []int{0, 2147483647, -2147483648, 127}
    if (1000 <= 0) {
        v4 = 300
    } else {
        switch ("foo bar" + "\n\t") {
        case "":
            v14 += v14
        case "a":
            if true {
                var v23 int8 = -1
                v24 := GenS4{F0: v4, F1: true, F2: "foo bar"}
                delete(v1, 1)
                _ = v23
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        _ = len(v22)
        _ = fi25
    }
    v26 := GenS2{F0: 127, F1: 386, F2: (v4 >> ((65535) & 63))}
    var v27 int8 = v14
    if (42 > 1) {
        fmt.Println(v27)
        var v28 int32 = -335
        _ = v28
    }
    if false {
        v29 := GenS5{F0: -977}
        for fi30 := 0; fi30 < 11; fi30++ {
            if (!false) {
                v31 := GenS2{F0: 2, F1: 4294967295, F2: 0}
                _ = v31
            } else {
                var v32 uint8 = 128
                v33 := v3[(fi30 / ((fi30) | 1))]
                _ = v32
                _ = v33
            }
            switch fi30 {
            case 747:
                _ = len(v22)
            }
            var v34 uint32 = 65535
            _ = fi30
            _ = v34
        }
        var v35 bool = (-880 == -32768)
        _ = v29
        _ = v35
    }
    if false {
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 uint16 = 1
            var v38 int64 = -1
            for fi39 := 0; fi39 < 2; fi39++ {
                switch (fi36 << (((v4 + (v4 % ((212) | 1)))) & 63)) {
                case 1:
                    var v40 int = fi39
                    _ = v40
                case 277:
                    var v41 uint32 = 977
                    _ = v41
                }
                _ = fi39
            }
            _ = fi36
            _ = v37
            _ = v38
        }
        var v42 bool = true
        _ = v42
    }
    switch 32768 {
    case 4294967295:
        fmt.Println(v4, v27, v4)
    }
    fmt.Println(v4, v14)
    v27 = (v27 | (^(v27)))
    var v43 uint32 = 4294967295
    switch 128 {
    case 128:
        var v44 uint8 = 127
        _ = v44
    case 4294967295:
        var v45 string = "\n\t"
        _ = v45
    case -760:
        var v46 bool = (!(32767 < 1000))
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v43
}
