package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0() uint {
    if (42 <= 222) {
        v7 := []int{2147483647, 4294967295, -1}
        if (18446744073709551615 >= 0) {
            var v8 int16 = 32767
            _ = v8
        } else {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint64 = 371
                _ = fi9
                _ = v10
            }
        }
        var v11 int = -1
        _ = v11
        _ = v7
    } else {
        var v12 uint16 = 32767
        v13 := []int{380, 32768, 4294967295, 794}
        var v14 uint = 2
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 string = "\n\t"
    v16 := []int{2147483647, 32768, 32768, -1}
    _ = v15
    _ = v16
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 0.1
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]uint32)
    var v5 string = "foo bar"
    var v6 int64 = -843
    var v17 bool = (v3 == v3)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 string = "a"
        switch v19 {
        case "\n\t":
            var v20 uint = 65535
            _ = v20
        }
        fmt.Println(fi18, v1)
        _ = fi18
        _ = v19
    }
    switch v5 {
    case "hello":
        fmt.Println(v1, v1, v3)
    default:
        var v21 float32 = (1e10 / 0.1)
        _ = v21
    }
    var v22 uint32 = v3
    v1 = (v1 ^ v1)
    switch -1 {
    case 32768:
        v6 = (v6 + 9223372036854775807)
    case 32767:
        for fi23 := 0; fi23 < 13; fi23++ {
            switch "a" {
            case "\n\t":
                switch (fi23 - fi23) {
                case 127:
                    switch v5 {
                    case "foo bar":
                        fmt.Println(v22, v6, fi23)
                    case "\n\t":
                        v3 -= (^(v22))
                    }
                case 9223372036854775807:
                    v24 := GenS1{F0: -771, F1: 32768}
                    _ = v24
                }
            }
            v1 -= (v1 + (v1 | (v1 | 2)))
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 uint = 2
                _ = fi25
                _ = v26
            }
            _ = fi23
        }
    default:
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float32 = (float32(v2) - float32(0.1))
            switch int(v1) {
            case -9223372036854775808:
                v29 := GenS1{F0: 0, F1: fi27}
                _ = v29
            case 1:
                v2 = 0.0
            }
            _ = fi27
            _ = v28
        }
    }
    _ = len(v4)
    var v30 uint8 = 42
    switch ("hello" + (v5 + v5)) {
    case "\n\t":
        v5 = v5
    default:
        if (-1 != int8(v3)) {
            var v31 int16 = 32767
            if v17 {
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 uint64 = 18446744073709551615
                    v34 := v4[uint32(v30)]
                    var v35 int16 = v31
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
                fmt.Println(v30, v3, v1)
                v36 := GenS1{F0: (v31 << ((272) & 63)), F1: 2147483647}
                _ = v36
            }
            v37 := GenS2{F0: 4294967295, F1: 21}
            _ = v31
            _ = v37
        } else {
            v38 := []int{-1, 32767, -2147483648}
            _ = v38
        }
    }
    v39 := []int{-929, int(v1), -1}
    v40 := GenS3{F0: 1, F1: (v6 % ((42) | 1)), F2: 15}
    v1 = (v1 & v1)
    switch ((v5 + v5) + ((v5 + v5) + (v5 + v5))) {
    case "hello":
        switch 2147483647 {
        case 339:
            var v41 int16 = 65
            _ = v41
        }
    }
    switch 42 {
    case 4294967295:
        v30 -= v30
    case 32767:
        v6 -= v6
    case 42:
        if v17 {
            v42 := GenS1{F0: 938, F1: 32767}
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
