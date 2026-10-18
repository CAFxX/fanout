package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 string = ""
    var v3 string = "\n\t"
    var v4 int = -750
    var v5 bool = true
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 8; fi7++ {
        switch v2 {
        case "hello":
            v5 = false
        default:
            v8 := GenS1{F0: 535, F1: 0}
            _ = v8
        }
        v2 = "a"
        for fi9 := 0; fi9 < 8; fi9++ {
            v2 = v3
            fmt.Println(v5, v4)
            _ = fi9
        }
        _ = fi7
    }
    var v10 bool = v5
    v4 += (v4 ^ v4)
    switch (((v2 + "") + "\n\t") + (((v2 + v2) + (v3 + "foo bar")) + ("\n\t" + v3))) {
    case "foo bar":
        switch v4 {
        case 495:
            var v11 float64 = 1e10
            _ = v11
        }
    case "":
        v12 := GenS1{F0: 93, F1: 0}
        _ = v12
    case "\n\t":
        if (!v5) {
            v6[""] = (4294967295 < 2147483648)
        }
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        switch "foo bar" {
        case "a":
            v5 = ((42 < -69) || (1000 != -781))
        case "foo bar":
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint = 1
                _ = fi14
                _ = v15
            }
        case "hello":
            fmt.Println(v4)
        }
        var v16 float64 = 1.5
        var v17 uint64 = 42
        _ = fi13
        _ = v16
        _ = v17
    }
    var v18 uint16 = 65535
    v10 = v10
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS2{F0: 236, F1: (^(v18)), F2: 240}
        v4 = fi19
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 int8 = 0
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    switch v4 {
    case 128:
        for fi23 := 0; fi23 < 12; fi23++ {
            v18 = (v18 << ((2) & 63))
            for fi24 := 0; fi24 < 17; fi24++ {
                v18 = (^(v18))
                v1[v2] = (fi24 == (fi24 * 682))
                if v5 {
                    if true {
                        var v25 int32 = -1
                        _ = len(v6)
                        _ = v25
                    } else {
                        v26 := GenS3{F0: -321}
                        var v27 string = v2
                        var v28 bool = true
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    v5 = (v10 || (1 >= 42))
                    v29 := GenS3{F0: -1}
                    _ = v29
                }
                fmt.Println(fi23, v3, v5)
                _ = fi24
            }
            _ = fi23
        }
    default:
        switch v2 {
        case "":
            v4 = v4
        case "foo bar":
            if true {
                if (!(971 < -257)) {
                    var v30 int64 = 0
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := []int{fi31}
                        _ = fi31
                        _ = v32
                    }
                    v4 = v4
                    _ = v30
                } else {
                    var v33 uint32 = 54
                    v33 = (v33 >> ((4294967295) & 63))
                    _ = v33
                }
                switch v3 {
                case "\n\t":
                    if (((1e10 / 0.1) + 3.14159) != -2.25) {
                        var v34 int32 = 2147483647
                        _ = v34
                    }
                }
                var v35 float64 = 0.0
                _ = v35
            }
        case "hello":
            if (42 < -2147483648) {
                switch v3 {
                case "":
                    var v36 float64 = (0.0 / (float64(float32(1.5)) + (float64((1e10 * 0.1)) * -2.25)))
                    _ = v36
                case "\n\t":
                    for fi37 := 0; fi37 < 10; fi37++ {
                        if (207 < 127) {
                            v10 = v5
                            switch v4 {
                            case 32768:
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    v39 := []int{}
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            var v40 float64 = (0.0 * 1.5)
                            _ = v40
                        }
                        fmt.Println(v3, v18, v3)
                        _ = fi37
                    }
                case "a":
                    v41 := GenS3{F0: 1}
                    _ = v41
                default:
                    v42 := GenS2{F0: -528, F1: (^(v18)), F2: 4294967295}
                    _ = v42
                }
            }
        default:
            fmt.Println(v18)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
