package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 uint64 = 4294967295
    var v4 uint16 = 42
    var v5 uint32 = 635
    var v6 uint = 397
    switch "" {
    case "":
        for fi7 := 0; fi7 < 4; fi7++ {
            if (0 <= (v5 + v5)) {
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 int8 = -1
                    _ = len(v1)
                    _ = fi8
                    _ = v9
                }
                var v10 bool = true
                _ = v10
            }
            var v11 float32 = 1e10
            var v12 int = (fi7 & 128)
            _ = fi7
            _ = v11
            _ = v12
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v4, v6, v5)
    v13 := GenS1{F0: (("hello" + "foo bar") + "hello"), F1: (v4 % ((v4) | 1)), F2: int32(v5)}
    var v14 uint8 = 255
    switch "foo bar" {
    case "foo bar":
        for fi15 := 0; fi15 < 12; fi15++ {
            v6 += v6
            if (!false) {
                if true {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        fmt.Println(v4, v4)
                        _ = fi16
                    }
                    v17 := GenS1{F0: ("" + ("" + ("hello" + "a"))), F1: (v4 * v4), F2: -1}
                    _ = v17
                }
                var v18 string = ((("a" + "hello") + "\n\t") + ("hello" + "hello"))
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 bool = false
                    v21 := GenS1{F0: v18, F1: v4, F2: 42}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := GenS1{F0: "a", F1: v4, F2: -2147483648}
                switch "\n\t" {
                case "":
                    v4 -= 347
                default:
                    var v24 string = "hello"
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            _ = fi15
        }
    case "":
        if true {
            if true {
                v5 = v5
                for fi25 := 0; fi25 < 11; fi25++ {
                    if false {
                        v6 = 867
                    }
                    _ = fi25
                }
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 float32 = -2.25
                    var v28 bool = false
                    v29 := GenS1{F0: "foo bar", F1: 0, F2: 42}
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            } else {
                var v30 bool = false
                if v30 {
                    var v31 uint32 = (^(v5))
                    v1[(v31 >> (((^(v6))) & 63))] = -1
                    _ = v31
                } else {
                    var v32 uint = v6
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v33 uint16 = v4
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 6; fi34++ {
                    v14 = 0
                    _ = len(v1)
                    _ = fi34
                }
                _ = v30
            }
            v3 = (^(v3))
        }
    }
    var v35 int32 = 1
    _ = len(v1)
    v3 = v3
    var v36 uint8 = v14
    var v37 int = -181
    v38 := GenS1{F0: "", F1: 1, F2: (-(v35))}
    if true {
        fmt.Println(v3, v14, v5)
        for fi39 := 0; fi39 < 15; fi39++ {
            v36 = (^(v36))
            _ = fi39
        }
        v40 := GenS2{F0: 0}
        _ = v40
    } else {
        fmt.Println(v36)
        v41 := GenS1{F0: "", F1: v4, F2: 0}
        _ = v41
    }
    v37 = v37
    var v42 uint64 = (v3 | v3)
    var v43 int64 = 9223372036854775807
    fmt.Println(v6, v4, v42)
    var v44 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
