package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 0
    var v3 uint = 4294967295
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{fi5, (fi5 * -142), fi5}
        for fi7 := 0; fi7 < 15; fi7++ {
            delete(v1, "hello")
            if false {
                var v8 float32 = 1.5
                var v9 uint16 = (^(v2))
                delete(v1, "hello")
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v3 += (v3 ^ v3)
    if false {
        v2 = (^(v2))
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int16 = -168
        v3 += v3
        _ = fi10
        _ = v11
    }
    v4 = (v4 % ((v4) | 1))
    v4 = 0
    fmt.Println(v4, v2)
    for fi12 := 0; fi12 < 18; fi12++ {
        if true {
            v3 += v3
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: (v3 & v3), F1: fi13, F2: 32767}
                fmt.Println(fi12)
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    if true {
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!true) {
                switch fi15 {
                case 127:
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS2{F0: ("a" + "\n\t"), F1: ""}
                        _ = fi16
                        _ = v17
                    }
                case -9223372036854775808:
                    if false {
                        switch fi15 {
                        case -338:
                            var v18 float64 = 1e10
                            _ = v18
                        case -2147483648:
                            var v19 string = "\n\t"
                            _ = v19
                        }
                        _ = len(v1)
                    } else {
                        var v20 uint32 = 2
                        if false {
                            var v21 float32 = (-2.25 + -2.25)
                            _ = v21
                        }
                        _ = v20
                    }
                case -15:
                    var v22 int16 = -1
                    _ = v22
                default:
                    v23 := v1["hello"]
                    _ = v23
                }
            } else {
                var v24 uint64 = 2
                v25 := v1["\n\t"]
                v26 := GenS3{F0: v24, F1: -1}
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 uint64 = 1
            _ = fi15
            _ = v27
        }
    }
    if (-2147483648 >= 0) {
        fmt.Println(v3, v4, v3)
        v28 := []int{32768, 1}
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
