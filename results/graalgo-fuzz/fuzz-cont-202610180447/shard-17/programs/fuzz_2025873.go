package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 int8, p1 int16) uint8 {
    fmt.Println(p1, p1)
    switch 127 {
    case -142:
        if (-2.25 == 0.0) {
            fmt.Println(p1, p0)
        }
    case 32768:
        var v5 float32 = (3.14159 * (1.5 * 3.14159))
        _ = v5
    case 1:
        if true {
            var v6 string = "hello"
            p0 -= p0
            _ = v6
        }
    default:
        if false {
            switch 1 {
            case -199:
                var v7 uint8 = 128
                _ = v7
            case -633:
                var v8 int32 = -930
                _ = v8
            default:
                var v9 int = -700
                _ = v9
            }
        } else {
            var v10 int16 = (-(p1))
            v11 := GenS1{F0: -32768}
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(v10, fi12)
                switch (fi12 * (fi12 << ((1) & 63))) {
                case 127:
                    fmt.Println(v10, fi12)
                case 714:
                    if false {
                        v13 := GenS1{F0: 1000}
                        v14 := []int{(fi12 ^ (fi12 << ((876) & 63)))}
                        _ = v13
                        _ = v14
                    } else {
                        switch -752 {
                        case -1:
                            v15 := GenS1{F0: p1}
                            _ = v15
                        }
                        if (1 > 18446744073709551615) {
                            v16 := []int{127, -2147483648}
                            var v17 uint32 = 65535
                            _ = v16
                            _ = v17
                        } else {
                            p1 += (^(p1))
                            v18 := []int{(fi12 / (((fi12 - fi12)) | 1)), fi12, (fi12 >> ((0) & 63))}
                            _ = v18
                        }
                        var v19 uint8 = 42
                        _ = v19
                    }
                default:
                    if true {
                        var v20 uint = 65535
                        v20 += v20
                        _ = v20
                    }
                }
                _ = fi12
            }
            _ = v10
            _ = v11
        }
    }
    var v21 uint16 = 1
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{fi22, fi22}
        if true {
            if false {
                for fi24 := 0; fi24 < 2; fi24++ {
                    fmt.Println(v21)
                    switch "" {
                    case "foo bar":
                        var v25 float64 = 1.5
                        _ = v25
                    }
                    _ = len(v23)
                    _ = fi24
                }
                if len(v23) > 0 {
                    v23[0] = 0
                }
            } else {
                v26 := GenS2{F0: fi22}
                if false {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        p1 = (p1 << ((18446744073709551615) & 63))
                        var v28 bool = (true && ((fi27 - (-(fi22))) < (fi27 % (((fi27 * (fi22 % ((-644) | 1)))) | 1))))
                        var v29 uint16 = v21
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                var v30 int8 = -83
                _ = v26
                _ = v30
            }
            p0 = (p0 - p0)
        }
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v21
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -779
    v3 := make(map[string]bool)
    var v4 int32 = int32(v2)
    if (v2 <= v2) {
        var v31 string = ""
        _ = v31
    } else {
        v2 = 1000
    }
    if (((((1.5 / 1e10) * ((3.14159 - 3.14159) / 1e10)) * -2.25) / (0.1 / ((-2.25 + (3.14159 + 0.1)) - 1e10))) < -2.25) {
        fmt.Println(v4, v2)
        v2 += v2
        var v32 uint = 42
        _ = v32
    } else {
        var v33 uint64 = 0
        _ = v33
    }
    v4 = 0
    v2 += (v2 - (v2 ^ 32767))
    switch -2147483648 {
    case -9223372036854775808:
        if false {
            fmt.Println(v2, v2)
            switch "foo bar" {
            case "a":
                if (("hello" < ("hello" + "\n\t")) && (float64(1.5) != -2.25)) {
                    v4 = 0
                }
            case "hello":
                switch ("\n\t" + ("a" + "")) {
                case "\n\t":
                    v34 := GenS3{F0: "foo bar"}
                    _ = v34
                case "hello":
                    v35 := []int{2147483647, 127}
                    _ = v35
                }
            case "\n\t":
                var v36 bool = true
                _ = v36
            }
        } else {
            var v37 uint = 4294967295
            _ = v37
        }
    case -493:
        fmt.Println(v4, v4)
    }
    var v38 int32 = (v4 << ((42) & 63))
    delete(v1, 318)
    v39 := GenS3{F0: ""}
    var v40 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
