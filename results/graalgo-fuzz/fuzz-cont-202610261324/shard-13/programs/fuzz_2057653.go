package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = (v1 << ((v1) & 63))
    v3 := make(map[int]bool)
    var v4 bool = true
    var v5 uint64 = 2
    var v6 uint8 = 103
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (fi7 | (-(fi7))) {
        case 128:
            for fi8 := 0; fi8 < 17; fi8++ {
                v6 = v6
                v2 += 18446744073709551615
                _ = fi8
            }
        default:
            var v9 int64 = 733
            _ = v9
        }
        v10 := GenS1{F0: -2147483648, F1: v6}
        _ = fi7
        _ = v10
    }
    fmt.Println(v1, v1)
    v6 += (^(v6))
    if true {
        v1 = 539
    }
    var v11 int64 = -9223372036854775808
    for fi12 := 0; fi12 < 5; fi12++ {
        switch "" {
        case "a":
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 uint = (v1 << ((255) & 63))
                v1 = v1
                _ = fi13
                _ = v14
            }
        default:
            v15 := GenS1{F0: 1, F1: 15}
            _ = v15
        }
        if (true && (-871 < 1000)) {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int8 = 100
                v17 -= v17
                fmt.Println(v11)
                _ = fi16
                _ = v17
            }
            if (!v4) {
                v18 := []int{(fi12 % ((-660) | 1))}
                v1 = v1
                v6 = (v6 / ((v6) | 1))
                _ = v18
            } else {
                switch (fi12 % (((fi12 & int(v6))) | 1)) {
                case -2147483648:
                    switch "\n\t" {
                    case "foo bar":
                        v19 := []int{fi12, 32768, (fi12 ^ fi12)}
                        _ = v19
                    case "hello":
                        v20 := GenS1{F0: -988, F1: 17}
                        _ = v20
                    case "\n\t":
                        delete(v3, fi12)
                    }
                case -318:
                    for fi21 := 0; fi21 < 1; fi21++ {
                        var v22 int32 = 443
                        _ = fi21
                        _ = v22
                    }
                case -955:
                    if (v4 && v4) {
                        if false {
                            v23 := GenS1{F0: -1, F1: 255}
                            _ = v23
                        }
                    } else {
                        switch "a" {
                        case "foo bar":
                            for fi24 := 0; fi24 < 20; fi24++ {
                                var v25 float64 = 3.14159
                                _ = fi24
                                _ = v25
                            }
                        default:
                            v26 := GenS2{F0: (v11 << ((2) & 63))}
                            _ = v26
                        }
                        switch fi12 {
                        case 531:
                            v3[128] = (((0.1 / 0.1) != -2.25) || ((fi12 ^ -9223372036854775808) > fi12))
                        case -199:
                            v11 = (v11 & v11)
                        default:
                            v11 = (v11 - (^(v11)))
                        }
                        switch "a" {
                        case "":
                            var v27 string = "foo bar"
                            _ = v27
                        case "hello":
                            fmt.Println(fi12, fi12)
                        default:
                            for fi28 := 0; fi28 < 20; fi28++ {
                                v5 -= v5
                                _ = fi28
                            }
                        }
                    }
                default:
                    v11 += v11
                }
                v29 := []int{(fi12 ^ fi12), fi12}
                _ = v29
            }
            if v4 {
                switch "" {
                case "a":
                    v4 = ((fi12 / (((^(fi12))) | 1)) == fi12)
                default:
                    var v30 int16 = 32767
                    _ = v30
                }
                switch "" {
                case "a":
                    v5 -= uint64(v2)
                case "hello":
                    v31 := []int{fi12, (fi12 >> ((v2) & 63))}
                    _ = v31
                case "":
                    v32 := []int{fi12}
                    _ = v32
                default:
                    var v33 float32 = ((float32(1.5) - ((1.5 * ((1e10 - float32(3.14159)) - 0.1)) + (1e10 - -2.25))) + (-2.25 / (float32((0.0 - -2.25)) - 1.5)))
                    _ = v33
                }
                var v34 int = fi12
                _ = v34
            }
        }
        _ = fi12
    }
    v35 := GenS3{F0: int64(v2)}
    var v36 string = "hello"
    v37 := GenS1{F0: -1, F1: (v6 / ((2) | 1))}
    var v38 string = v36
    var v39 int = 448
    v40 := GenS4{F0: 191}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
