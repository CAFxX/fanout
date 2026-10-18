package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int8) uint16 {
    p0 = p0
    var v7 int8 = (p0 ^ 77)
    var v8 uint32 = 65535
    v9 := GenS1{F0: -1, F1: 32767}
    if true {
        v10 := []int{128, -899}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if (1 <= 0) {
                var v12 int64 = 0
                var v13 uint = 0
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                v7 += p0
                _ = fi14
            }
            _ = fi11
        }
        var v15 uint32 = v8
        _ = v15
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 1
    var v2 int = 4294967295
    var v3 int16 = 1
    v4 := []string{"\n\t"}
    v5 := make(map[string]int)
    var v6 uint32 = 2147483647
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint = 422
        switch (fi16 + 437) {
        case 1:
            v5[(("\n\t" + "foo bar") + ("\n\t" + "foo bar"))] = 127
        case 128:
            switch fi16 {
            case -572:
                v1 += v1
            }
        }
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v1 += (v1 | (v1 ^ 0))
        var v19 int32 = -1
        var v20 float64 = (1.5 - 0.0)
        _ = fi18
        _ = v19
        _ = v20
    }
    if false {
        if true {
            if ("\n\t" != ("\n\t" + "foo bar")) {
                var v21 uint64 = 0
                for fi22 := 0; fi22 < 16; fi22++ {
                    v23 := GenS2{F0: 65535, F1: -1, F2: 178}
                    fmt.Println(v6)
                    if true {
                        v3 = (v3 & (v3 | v3))
                        v2 -= (v2 ^ (fi22 / ((42) | 1)))
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v25 := v5["hello"]
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                v3 = (v3 * -877)
                if (1 < 0) {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        if false {
                            var v27 int32 = -1
                            _ = v27
                        }
                        if (2 >= 2) {
                            var v28 string = ""
                            switch v28 {
                            case "hello":
                                for fi29 := 0; fi29 < 8; fi29++ {
                                    v2 = fi29
                                    _ = fi29
                                }
                            }
                            _ = v28
                        }
                        _ = fi26
                    }
                    switch (("" + "a") + ("hello" + "hello")) {
                    case "a":
                        var v30 float32 = 1e10
                        _ = v30
                    }
                    v1 += (v1 % ((v1) | 1))
                } else {
                    var v31 uint8 = 255
                    if true {
                        var v32 float64 = 0.0
                        v32 -= v32
                        v2 -= 4294967295
                        _ = v32
                    } else {
                        var v33 bool = false
                        _ = len(v4)
                        v1 = v1
                        _ = v33
                    }
                    if true {
                        fmt.Println(v2, v1, v2)
                        var v34 uint64 = 852
                        v35 := GenS3{F0: (v2 | v2), F1: 42}
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int8 = 127
                        _ = v36
                    }
                    _ = v31
                }
            }
            var v37 uint32 = (v6 >> ((42) & 63))
            _ = v37
        }
        fmt.Println(v1)
        var v38 string = ("a" + "")
        _ = v38
    }
    if false {
        var v39 uint8 = 213
        if true {
            var v40 int64 = 0
            v40 -= 42
            _ = v40
        }
        v3 += v3
        _ = v39
    }
    var v41 int32 = 273
    v42 := GenS1{F0: (v41 - -979), F1: 366}
    var v43 int64 = -1
    fmt.Println(v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
