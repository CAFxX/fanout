package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = -2.25
    var v3 int64 = 42
    fmt.Println(v2, v3)
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float64 = 3.14159
        for fi6 := 0; fi6 < 16; fi6++ {
            switch (^(fi4)) {
            case 32767:
                var v7 int32 = 2147483647
                _ = v7
            }
            if ((v2 / (float32(v5) - 1.5)) != -2.25) {
                var v8 uint8 = 160
                var v9 int8 = 1
                var v10 uint8 = (v8 - v8)
                _ = v10
                _ = v8
                _ = v9
            } else {
                _ = len(v1)
                fmt.Println(v2, v2, v5)
                fmt.Println(v3, fi4, v3)
            }
            switch ("a" + ("a" + "a")) {
            case "a":
                var v11 string = "hello"
                _ = v11
            default:
                var v12 uint = 4294967295
                _ = v12
            }
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    var v13 uint8 = 134
    if (false || (!(v3 > (v3 << ((65535) & 63))))) {
        if true {
            v2 += v2
            var v14 uint8 = 0
            fmt.Println(v14, v14, v2)
            _ = v14
        } else {
            if false {
                v13 = 204
                for fi15 := 0; fi15 < 5; fi15++ {
                    if false {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            var v17 uint64 = 42
                            if false {
                                if (-110 <= -85) {
                                    v18 := []int{fi16, 127, (fi16 | 42)}
                                    var v19 uint64 = 316
                                    _ = v18
                                    _ = v19
                                }
                                v2 = v2
                            } else {
                                v17 = (v17 & (v17 >> ((1) & 63)))
                                v1["\n\t"] = int(v17)
                            }
                            fmt.Println(v3, fi15)
                            v3 = v3
                            _ = fi16
                            _ = v17
                        }
                        v20 := GenS1{F0: fi15}
                        _ = v20
                    }
                    _ = fi15
                }
            }
            fmt.Println(v3)
            switch ("\n\t" + "\n\t") {
            case "hello":
                if (!(-1 > 128)) {
                    delete(v1, ("" + "hello"))
                }
            case "\n\t":
                v21 := []int{-253}
                _ = v21
            case "a":
                var v22 float64 = 1e10
                _ = v22
            }
        }
        if (128 < (^(v13))) {
            switch (("" + ("foo bar" + "\n\t")) + ("a" + ((("hello" + ("\n\t" + ("" + "hello"))) + "\n\t") + "hello"))) {
            case "\n\t":
                v3 += (^(v3))
            case "hello":
                if false {
                    v23 := GenS2{F0: 2, F1: 1}
                    v24 := GenS3{F0: v13}
                    if true {
                        var v25 int64 = (v3 + v3)
                        switch ((("hello" + "") + "") + "a") {
                        case "foo bar":
                            v13 = (v13 ^ 42)
                        case "hello":
                            v3 = v25
                        }
                        _ = v25
                    } else {
                        var v26 float64 = 3.14159
                        _ = v26
                    }
                    _ = v23
                    _ = v24
                } else {
                    v27 := GenS4{F0: 584, F1: (^(v13)), F2: uint32(v13)}
                    fmt.Println(v2, v3)
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v29 := GenS4{F0: fi28, F1: 21, F2: 42}
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
            }
            if (!true) {
                if false {
                    v2 = float32(3.14159)
                    var v30 uint32 = 2147483647
                    _ = v30
                }
            }
            v2 -= 0.0
        } else {
            v31 := []int{}
            var v32 float64 = 0.0
            var v33 int16 = 32767
            _ = v31
            _ = v32
            _ = v33
        }
    } else {
        v34 := v1[""]
        fmt.Println(v2, v2, v3)
        _ = v34
    }
    v13 = v13
    var v35 string = "hello"
    var v36 bool = false
    v2 += ((1e10 + v2) / (v2 + (((float32(0.1) / 1.5) - v2) * v2)))
    for fi37 := 0; fi37 < 18; fi37++ {
        for fi38 := 0; fi38 < 4; fi38++ {
            v35 = v35
            var v39 uint32 = 42
            _ = fi38
            _ = v39
        }
        if v36 {
            v35 = ("\n\t" + v35)
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            fmt.Println(v36)
            _ = fi40
        }
        _ = fi37
    }
    var v41 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v41
}
