package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int16 }
type GenS3 struct { F0 int32 }
func genFn0() int32 {
    v5 := []int{385, 1, 720}
    var v6 int64 = 96
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 = (v6 << ((646) & 63))
        fmt.Println(fi7, v6)
        v5 = append(v5, fi7)
        _ = fi7
    }
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 2147483647
    v3 := []string{}
    v4 := make(map[uint32]int)
    v8 := v1[4294967295]
    fmt.Println(v2, v8)
    var v9 uint16 = 178
    switch (v2 >> ((0) & 63)) {
    case 32768:
        for fi10 := 0; fi10 < 6; fi10++ {
            if true {
                v8 -= 535
            }
            if ((v2 != 127) && (!true)) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 bool = (!((!true) || (!(9223372036854775807 != 2147483647))))
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi10
        }
    }
    v4[0] = v2
    if (128 <= 2) {
        v9 = (v9 + (v9 ^ (v9 - v9)))
        v9 += (v9 | v9)
        for fi13 := 0; fi13 < 17; fi13++ {
            if true {
                var v14 int8 = 127
                var v15 uint = 999
                var v16 uint64 = 4294967295
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 1000}
            _ = fi13
            _ = v17
        }
    }
    v18 := GenS2{F0: 2147483648, F1: genFn0(), F2: 0}
    var v19 uint = 2
    if true {
        v20 := GenS1{F0: -99}
        _ = v20
    } else {
        v21 := GenS1{F0: 42}
        for fi22 := 0; fi22 < 4; fi22++ {
            if (!false) {
                var v23 uint32 = 642
                _ = v23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 int8 = 8
                v25 -= (v25 & (v25 % ((-128) | 1)))
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                v27 := v1[249]
                v28 := GenS2{F0: 1, F1: 42, F2: 42}
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi22
        }
        v8 = (v8 ^ 2147483647)
        _ = v21
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v19, v2, v9)
    default:
        var v29 string = ("" + "\n\t")
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        if false {
            if (!false) {
                v2 -= (v2 ^ v8)
                v3 = append(v3, "foo bar")
            } else {
                var v31 bool = (1 != (^(v9)))
                _ = v31
            }
            switch v2 {
            case 227:
                v32 := GenS1{F0: 32767}
                _ = v32
            }
            switch (v8 | v2) {
            case 462:
                v33 := GenS3{F0: -735}
                _ = v33
            default:
                switch "hello" {
                case "a":
                    v34 := GenS1{F0: 42}
                    _ = v34
                case "":
                    v35 := GenS3{F0: -916}
                    _ = v35
                case "hello":
                    fmt.Println(v19, fi30, v8)
                }
            }
        } else {
            var v36 uint16 = v9
            for fi37 := 0; fi37 < 20; fi37++ {
                fmt.Println(fi30, v19)
                if false {
                    if (true || false) {
                        var v38 int64 = 202
                        _ = v38
                    } else {
                        switch -204 {
                        case -9223372036854775808:
                            var v39 string = "a"
                            _ = v39
                        default:
                            if (((0.1 + 0.0) >= ((0.1 / 0.1) - float32(1.5))) && true) {
                                fmt.Println(v9, v36, v9)
                            } else {
                                v4[4294967295] = (fi30 & (fi30 & 452))
                            }
                        }
                        var v40 int16 = 178
                        _ = v40
                    }
                    fmt.Println(v8, v36)
                } else {
                    var v41 int64 = 9223372036854775807
                    _ = v41
                }
                _ = fi37
            }
            _ = v36
        }
        _ = fi30
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        for fi43 := 0; fi43 < 20; fi43++ {
            v2 += v8
            _ = fi43
        }
        v1[463] = fi42
        _ = fi42
    }
    var v44 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v8
    _ = v9
}
