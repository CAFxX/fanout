package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 float32 = 3.14159
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int32 = -964
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(v7, v3)
            v9 := GenS1{F0: ("\n\t" + (("hello" + ("\n\t" + "\n\t")) + ("" + "")))}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v5[821] = 32768
    switch "" {
    case "":
        if (1e10 != 0.0) {
            v3 -= v3
        }
    case "hello":
        if (!((("" + "a") + "foo bar") == "")) {
            var v10 uint8 = 108
            if false {
                var v11 int = 127
                fmt.Println(v10)
                for fi12 := 0; fi12 < 12; fi12++ {
                    if true {
                        v11 = fi12
                    } else {
                        v3 = v3
                        v13 := GenS1{F0: ("\n\t" + "")}
                        var v14 int64 = 981
                        _ = v13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := []int{v11, (v11 | v11)}
                        if len(v16) > 0 {
                            v16[0] = 944
                        }
                        v16 = append(v16, v11)
                        _ = fi15
                        _ = v16
                    }
                    _ = fi12
                }
                _ = v11
            }
            _ = v10
        } else {
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS1{F0: ("\n\t" + ("foo bar" + "a"))}
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := []int{(fi19 >> ((535) & 63)), (-(fi17)), (fi17 * (fi17 & (fi17 * (-(fi19))))), 809}
                    _ = fi19
                    _ = v20
                }
                v3 = float32((1.5 * 0.1))
                _ = fi17
                _ = v18
            }
            v21 := []int{128, 1}
            var v22 int8 = 100
            _ = v21
            _ = v22
        }
    case "a":
        for fi23 := 0; fi23 < 2; fi23++ {
            switch "a" {
            case "hello":
                switch ("" + ("hello" + "\n\t")) {
                case "foo bar":
                    v24 := GenS2{F0: fi23}
                    _ = v24
                case "hello":
                    if (0 >= 65535) {
                        v3 = v3
                        v25 := []int{fi23, fi23, fi23}
                        for fi26 := 0; fi26 < 3; fi26++ {
                            var v27 uint = 676
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    }
                default:
                    var v28 int8 = 127
                    _ = v28
                }
            case "\n\t":
                for fi29 := 0; fi29 < 4; fi29++ {
                    fmt.Println(fi23, v3, fi23)
                    v30 := []int{(fi23 / (((fi29 ^ fi29)) | 1))}
                    _ = fi29
                    _ = v30
                }
            }
            switch "\n\t" {
            case "":
                v3 = v3
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                v3 = (v3 + (v3 + 0.1))
                var v32 float32 = 3.14159
                _ = fi31
                _ = v32
            }
            _ = fi23
        }
    }
    if (-32768 == -981) {
        v3 -= (3.14159 * (((v3 * 1.5) * float32(((-2.25 - 1.5) + 3.14159))) * 1.5))
        if true {
            for fi33 := 0; fi33 < 15; fi33++ {
                fmt.Println(fi33)
                _ = fi33
            }
            if false {
                v3 -= v3
            }
            var v34 bool = true
            _ = v34
        }
        var v35 uint64 = 47
        _ = v35
    }
    v3 = v3
    if ((0 == 42) || (997 <= 796)) {
        v3 = (v3 - v3)
    } else {
        v3 -= v3
    }
    var v36 uint8 = 1
    delete(v2, 4294967295)
    v3 = (3.14159 * 0.1)
    v37 := GenS2{F0: 1}
    var v38 uint8 = v36
    var v39 int32 = -2147483648
    var v40 int64 = 42
    if (38 >= 0) {
        var v41 uint = 255
        _ = v41
    }
    v40 += v40
    var v42 int64 = -307
    if (true || (v38 <= (v36 >> ((75) & 63)))) {
        if ((4294967295 == 0) && (float64(v3) != (-2.25 * ((float64(v3) + 0.0) * 0.1)))) {
            v40 = (v40 ^ v42)
            fmt.Println(v39, v42)
        } else {
            v43 := GenS2{F0: 592}
            _ = v43
        }
    } else {
        v42 -= v40
        v3 = (v3 / (v3 - float32(((1.5 / -2.25) + float64(v3)))))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
