package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 128
    var v3 int32 = -375
    if true {
        for fi4 := 0; fi4 < 18; fi4++ {
            for fi5 := 0; fi5 < 3; fi5++ {
                if true {
                    v3 -= 2147483647
                    fmt.Println(v2)
                    v6 := GenS1{F0: -128, F1: (fi5 & (fi5 ^ -9223372036854775808)), F2: fi5}
                    _ = v6
                }
                if true {
                    v3 -= (v3 >> ((42) & 63))
                    if (0 >= 76) {
                        v7 := []int{1, (fi4 ^ (^(fi4))), fi4}
                        v3 = 780
                        _ = len(v7)
                        _ = v7
                    } else {
                        v2 -= (v2 ^ 1)
                    }
                }
                fmt.Println(v3, fi5)
                _ = fi5
            }
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS2{F0: 9223372036854775807}
                _ = fi8
                _ = v9
            }
            if true {
                v3 = -117
            }
            _ = fi4
        }
        v10 := GenS2{F0: 0}
        switch -8 {
        case 500:
            for fi11 := 0; fi11 < 5; fi11++ {
                switch (("" + ("hello" + "a")) + "a") {
                case "hello":
                    v12 := []int{}
                    _ = v12
                }
                _ = fi11
            }
        case 128:
            var v13 int64 = 436
            _ = v13
        }
        _ = v10
    }
    switch "foo bar" {
    case "foo bar":
        var v14 float64 = 3.14159
        _ = v14
    case "hello":
        if (!true) {
            var v15 bool = true
            _ = v15
        }
    case "a":
        if (!true) {
            var v16 float64 = (0.0 + 0.1)
            v17 := GenS3{F0: -1, F1: (v1 ^ v1), F2: 255}
            _ = v16
            _ = v17
        } else {
            v18 := []int{0}
            v2 += 0
            _ = v18
        }
    }
    if true {
        v19 := GenS4{F0: 620, F1: 42, F2: 0}
        var v20 float32 = (3.14159 - (-2.25 / (float32(1e10) - float32(-2.25))))
        var v21 float32 = 1e10
        _ = v19
        _ = v20
        _ = v21
    } else {
        fmt.Println(v3, v3)
    }
    v22 := GenS2{F0: 9223372036854775807}
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            if false {
                var v25 int64 = -9223372036854775808
                _ = v25
            }
            _ = fi24
        }
        fmt.Println(v2, v1, fi23)
        var v26 uint = 42
        _ = fi23
        _ = v26
    }
    v27 := []int{}
    v27 = append(v27, 32767)
    var v28 int64 = 802
    if (!(true || true)) {
        v3 -= 2147483647
        v1 += v1
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 string = ("\n\t" + (("foo bar" + "a") + ("" + "foo bar")))
            _ = fi29
            _ = v30
        }
    } else {
        var v31 uint8 = (v2 + 255)
        _ = v31
    }
    v32 := GenS2{F0: (v28 ^ 0)}
    switch "" {
    case "hello":
        v33 := GenS5{F0: 100}
        _ = v33
    case "foo bar":
        if len(v27) > 0 {
            v27[0] = 128
        }
    case "a":
        var v34 uint16 = 255
        _ = v34
    }
    fmt.Println(v2)
    var v35 uint64 = v1
    var v36 uint64 = v35
    v27 = append(v27, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v35
    _ = v36
}
