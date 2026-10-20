package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int64 = -556
    v2 := make(map[string]int)
    var v3 uint16 = 65535
    var v4 int8 = 0
    v4 -= (v4 - v4)
    var v5 int8 = (v4 * 14)
    v5 = v5
    for fi6 := 0; fi6 < 3; fi6++ {
        if ((1 <= 0) || true) {
            switch 1 {
            case -11:
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 int16 = 32767
                    var v9 bool = ("a" > "hello")
                    v1 = (v1 + v1)
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 128:
                var v10 int8 = (v5 ^ (v4 | v4))
                _ = v10
            case -2147483648:
                for fi11 := 0; fi11 < 6; fi11++ {
                    v4 = v4
                    fmt.Println(v5)
                    v12 := GenS1{F0: 100}
                    _ = fi11
                    _ = v12
                }
            }
            if false {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 int32 = 887
                    _ = fi13
                    _ = v14
                }
                v15 := v2["\n\t"]
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                v4 = (v5 | v5)
                var v17 int32 = 1
                v18 := []int{0, 0}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v19 := GenS2{F0: -114}
        _ = fi6
        _ = v19
    }
    var v20 int = 565
    _ = len(v2)
    var v21 int64 = -1
    v2["\n\t"] = 1
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS2{F0: v5}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 65535, F1: (^(v3))}
    fmt.Println(v5, v5, v1)
    if true {
        fmt.Println(v3, v20, v1)
    } else {
        if (v3 == (v3 & (v3 + 73))) {
            if (true || true) {
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 float64 = 3.14159
                    v26 -= (float64(1.5) - (v26 - ((v26 / 1.5) + v26)))
                    switch (v20 << ((402) & 63)) {
                    case 9223372036854775807:
                        v27 := GenS2{F0: -75}
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
            } else {
                for fi28 := 0; fi28 < 4; fi28++ {
                    fmt.Println(v1, v21, v3)
                    for fi29 := 0; fi29 < 17; fi29++ {
                        if true {
                            v5 = (v4 % ((int8(v3)) | 1))
                            if false {
                                var v30 uint8 = 0
                                var v31 uint = 604
                                v20 -= -416
                                _ = v30
                                _ = v31
                            }
                        } else {
                            var v32 bool = false
                            _ = v32
                        }
                        if (42 > 127) {
                            if ((32767 > -1) || true) {
                                v33 := GenS1{F0: v4}
                                fmt.Println(fi29)
                                var v34 float32 = -2.25
                                _ = v33
                                _ = v34
                            } else {
                                v20 = fi28
                                v21 += v1
                                var v35 string = "\n\t"
                                _ = v35
                            }
                            v36 := GenS3{F0: 830, F1: 32767}
                            _ = v36
                        } else {
                            var v37 int64 = (v1 ^ v21)
                            v20 = 32768
                            _ = v37
                        }
                        _ = fi29
                    }
                    _ = fi28
                }
                fmt.Println(v21)
            }
            var v38 string = "foo bar"
            _ = v38
        } else {
            var v39 string = "\n\t"
            v40 := GenS2{F0: 42}
            fmt.Println(v20)
            _ = v39
            _ = v40
        }
        switch 128 {
        case -1:
            v41 := GenS2{F0: v4}
            _ = v41
        case 1:
            fmt.Println(v21, v20)
        }
    }
    fmt.Println(v20, v3, v1)
    v5 -= 103
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
