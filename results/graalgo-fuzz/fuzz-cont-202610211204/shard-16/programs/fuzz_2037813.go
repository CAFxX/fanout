package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 int64 }
func genFn0() int {
    var v5 uint8 = 42
    var v6 float32 = (-2.25 * 3.14159)
    _ = v5
    _ = v6
    return 487
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]string)
    var v3 uint8 = 30
    v4 := make(map[string]int)
    if (v3 > (v3 % ((42) | 1))) {
        var v7 int8 = -41
        v8 := GenS1{F0: 65535, F1: -1, F2: 97}
        v9 := GenS1{F0: 2, F1: -1, F2: -1}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if false {
            var v11 uint = 18446744073709551615
            _ = v11
        }
        switch "foo bar" {
        case "a":
            var v12 int16 = 32767
            _ = v12
        }
        _ = fi10
    }
    var v13 uint8 = v3
    if false {
        if false {
            v3 = 127
            switch "\n\t" {
            case "":
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            }
            v14 := GenS2{F0: 129, F1: "", F2: 32767}
            _ = v14
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint = 2
            var v17 uint32 = 42
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        fmt.Println(v13)
        v1 = append(v1, 2147483648)
        v13 = v13
    }
    var v18 uint = 4294967295
    var v19 uint32 = 2147483647
    if ((46 == 42) && false) {
        v19 -= (v19 & uint32(v3))
    }
    var v20 int = 127
    _ = len(v1)
    for fi21 := 0; fi21 < 2; fi21++ {
        if true {
            switch (fi21 | v20) {
            case -329:
                var v22 float64 = (float64(-2.25) / 3.14159)
                _ = v22
            case 262:
                for fi23 := 0; fi23 < 6; fi23++ {
                    v3 = v3
                    _ = fi23
                }
            }
        } else {
            var v24 int64 = 0
            v25 := GenS1{F0: 880, F1: 831, F2: 127}
            if (0 == -1) {
                var v26 string = "hello"
                v20 = (v20 * int(v18))
                if (42 > uint64(v18)) {
                    v2[int(v18)] = (v26 + v26)
                    var v27 uint64 = 895
                    _ = v27
                }
                _ = v26
            } else {
                if ((v18 - uint(v3)) <= v18) {
                    v19 = (v19 & v19)
                } else {
                    v18 = v18
                    for fi28 := 0; fi28 < 8; fi28++ {
                        v19 = (v19 & v19)
                        var v29 int64 = int64(fi28)
                        fmt.Println(v20, v29)
                        fmt.Println(v24)
                        _ = fi28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 2; fi30++ {
                        var v31 uint16 = 32767
                        _ = fi30
                        _ = v31
                    }
                }
                var v32 uint16 = 455
                v3 += (^(v3))
                _ = v32
            }
            _ = v24
            _ = v25
        }
        switch ("foo bar" + "hello") {
        case "a":
            v13 = (v3 * 255)
        case "\n\t":
            var v33 uint16 = 32767
            _ = v33
        }
        _ = fi21
    }
    if ((-1 == 0) || true) {
        for fi34 := 0; fi34 < 7; fi34++ {
            switch v20 {
            case 127:
                var v35 uint = uint(v19)
                _ = v35
            }
            v19 = v19
            v3 = v3
            _ = fi34
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := GenS3{F0: -836, F1: (^(v20)), F2: -696}
            _ = fi36
            _ = v37
        }
    } else {
        var v38 float32 = 0.1
        _ = v38
    }
    var v39 int32 = 42
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
}
