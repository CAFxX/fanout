package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 int64, p1 int16) int64 {
    p0 -= p0
    var v7 int = 153
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1() uint32 {
    var v8 int16 = 549
    v8 = v8
    var v9 uint8 = 42
    fmt.Println(v9, v8)
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]uint32)
    v3 := []int{-9223372036854775808}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v6 := []string{}
    if ((((1e10 / 1.0) - 3.14159) / (3.14159 * 0.1)) > float64((0.1 / (1e10 - (-2.25 * 1.5))))) {
        delete(v4, 127)
        v10 := GenS1{F0: 353, F1: 42}
        _ = v10
    }
    v1 = v1
    fmt.Println(v1)
    fmt.Println(v1)
    v1 -= v1
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            v12 := GenS1{F0: fi11, F1: -128}
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int64 = 42
                if (float64(-2.25) > 1.5) {
                    _ = len(v2)
                    var v15 string = ""
                    var v16 int8 = 127
                    _ = v15
                    _ = v16
                } else {
                    if false {
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            var v17 string = (("foo bar" + ("foo bar" + ("\n\t" + "a"))) + "a")
                            v17 = v17
                            _ = v17
                        }
                    }
                }
                var v18 int64 = genFn0((-(v14)), 1)
                _ = fi13
                _ = v14
                _ = v18
            }
            _ = v12
        }
        if (-128 > 68) {
            switch (-(fi11)) {
            case -2147483648:
                if true {
                    fmt.Println(v1, v1)
                } else {
                    if false {
                        var v19 uint16 = 941
                        _ = v19
                    }
                    var v20 uint8 = 113
                    _ = v20
                }
            default:
                if (!false) {
                    v21 := GenS1{F0: fi11, F1: 127}
                    var v22 bool = false
                    _ = v21
                    _ = v22
                }
            }
        }
        _ = fi11
    }
    switch "foo bar" {
    case "\n\t":
        for fi23 := 0; fi23 < 5; fi23++ {
            switch fi23 {
            case 32768:
                v1 = genFn1()
            case -9223372036854775808:
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := GenS2{F0: 4294967295, F1: 1, F2: 959}
                    _ = len(v6)
                    var v26 int8 = 42
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            _ = fi23
        }
    default:
        _ = len(v4)
    }
    if false {
        if (!true) {
            var v27 string = ("\n\t" + "a")
            v28 := GenS3{F0: 494, F1: true, F2: 1}
            v1 -= v1
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            switch (fi29 - (fi29 % (((fi29 | 127)) | 1))) {
            case -268:
                var v30 int16 = -597
                _ = v30
            }
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 uint8 = 200
                _ = fi31
                _ = v32
            }
            var v33 uint16 = 65535
            _ = fi29
            _ = v33
        }
    } else {
        if (42 <= -621) {
            v1 = v1
            var v34 uint64 = 656
            fmt.Println(v1, v1)
            _ = v34
        }
        var v35 int64 = 42
        _ = v35
    }
    var v36 uint64 = 0
    v37 := GenS3{F0: 32767, F1: (0 != 127), F2: 409}
    v36 = (v36 >> ((11) & 63))
    if (-30 <= 42) {
        var v38 int16 = 459
        for fi39 := 0; fi39 < 18; fi39++ {
            if true {
                fmt.Println(v38)
                v38 = -45
                switch (fi39 >> ((42) & 63)) {
                case 0:
                    fmt.Println(v36)
                }
            }
            if len(v3) > 0 {
                v3[0] = (fi39 - fi39)
            }
            fmt.Println(fi39, v1)
            _ = fi39
        }
        _ = v38
    } else {
        var v40 float64 = 0.0
        v41 := GenS2{F0: 936, F1: 115, F2: 2}
        _ = v40
        _ = v41
    }
    v1 = genFn1()
    var v42 uint64 = 0
    if (v36 == v42) {
        v5[65535] = -1
    } else {
        var v43 float32 = (float32(3.14159) + 0.1)
        v3 = append(v3, -677)
        _ = v43
    }
    var v44 uint64 = 42
    v42 -= (v42 + (v44 % ((2) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
