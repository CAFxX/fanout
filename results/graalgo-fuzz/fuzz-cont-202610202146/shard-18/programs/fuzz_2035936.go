package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 int = 4294967295
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 9; fi5++ {
        v1 = v1
        _ = fi5
    }
    if v1 {
        v1 = v1
        if ((v3 + v3) > v3) {
            var v6 int64 = 9223372036854775807
            if true {
                for fi7 := 0; fi7 < 11; fi7++ {
                    v3 += v3
                    v3 = (v3 & v3)
                    v8 := GenS1{F0: 42}
                    _ = fi7
                    _ = v8
                }
            }
            _ = v6
        }
    } else {
        var v9 uint = 4294967295
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int64 = int64(fi10)
            _ = fi10
            _ = v11
        }
        v2[((("a" + "\n\t") + "foo bar") + "a")] = v3
        _ = v9
    }
    if (v1 && (89 == -23)) {
        if (!(100 > 127)) {
            if (76 <= 1) {
                switch v3 {
                case 2147483647:
                    fmt.Println(v1)
                case 547:
                    var v12 uint = 255
                    _ = v12
                case 0:
                    v3 += v3
                default:
                    var v13 int16 = 32767
                    _ = v13
                }
                var v14 uint16 = 28
                _ = v14
            } else {
                var v15 uint32 = 2147483648
                if v1 {
                    v16 := GenS1{F0: 18446744073709551615}
                    v3 += (v3 - (v3 % ((-9223372036854775808) | 1)))
                    v17 := GenS1{F0: 0}
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS2{F0: -128}
                    _ = v18
                }
                v19 := v2[""]
                _ = v15
                _ = v19
            }
        }
        if (0 == 18446744073709551615) {
            var v20 bool = v1
            if (42 > 872) {
                var v21 uint16 = 1
                v22 := GenS1{F0: 458}
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 17; fi23++ {
                v20 = (("foo bar" + "hello") >= "\n\t")
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 int8 = 127
                    v3 = 32767
                    v26 := v4[4294967295]
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                if false {
                    v27 := GenS1{F0: 1}
                    _ = v27
                }
                _ = fi23
            }
            _ = v20
        } else {
            v28 := GenS3{F0: 729}
            fmt.Println(v3, v1)
            var v29 string = (("a" + ("a" + ("" + "hello"))) + (("" + "hello") + "\n\t"))
            _ = v28
            _ = v29
        }
        fmt.Println(v3)
    } else {
        v30 := []int{151, (^(v3))}
        for fi31 := 0; fi31 < 17; fi31++ {
            for fi32 := 0; fi32 < 2; fi32++ {
                _ = len(v30)
                _ = fi32
            }
            _ = fi31
        }
        _ = v30
    }
    switch v3 {
    case -998:
        v1 = v1
    default:
        v33 := GenS4{F0: 201, F1: 42, F2: 828}
        _ = v33
    }
    var v34 int = (v3 + (v3 + v3))
    var v35 bool = v1
    fmt.Println(v34, v1, v3)
    v4[211] = 0
    _ = len(v4)
    var v36 uint = 0
    fmt.Println(v3, v36, v35)
    var v37 uint = v36
    if (!false) {
        v3 += 654
    } else {
        v38 := []int{}
        _ = v38
    }
    var v39 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
}
