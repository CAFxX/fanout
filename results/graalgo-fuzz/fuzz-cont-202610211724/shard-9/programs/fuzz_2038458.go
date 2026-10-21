package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint64 }
func genFn0(p0 int8, p1 int64) int16 {
    switch ("a" + ("a" + "\n\t")) {
    case "a":
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(p1, fi6, p0)
            _ = fi6
        }
    default:
        var v7 float64 = 1e10
        _ = v7
    }
    fmt.Println(p1, p1)
    p1 = 42
    p1 = -229
    _ = p0
    _ = p1
    return 1000
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 143
    var v3 int = 32767
    var v4 int16 = -32768
    var v5 string = ("" + "hello")
    v2 = v2
    v8 := GenS1{F0: 42, F1: true, F2: v3}
    var v9 string = v5
    for fi10 := 0; fi10 < 6; fi10++ {
        if true {
            v11 := []int{(v3 | (v3 % (((^(fi10))) | 1)))}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = v11
        }
        if true {
            var v12 uint = 428
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{fi13, fi13}
                if false {
                    switch v9 {
                    case "\n\t":
                        v9 = v9
                    case "hello":
                        switch v9 {
                        case "foo bar":
                            var v15 string = ((v5 + ("" + v9)) + ("foo bar" + "hello"))
                            _ = v15
                        }
                    }
                } else {
                    var v16 uint = v12
                    if true {
                        var v17 int32 = 42
                        _ = v17
                    }
                    _ = v16
                }
                v9 = (v5 + v5)
                _ = fi13
                _ = v14
            }
            var v18 int8 = 29
            _ = v12
            _ = v18
        } else {
            switch v9 {
            case "":
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 float32 = (-2.25 + ((0.1 - ((1e10 - float32(1e10)) + 1.5)) + 0.1))
                    v5 = v5
                    if true {
                        v21 := []int{-187}
                        var v22 int8 = 42
                        delete(v1, -1)
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
            case "hello":
                v23 := []int{}
                _ = v23
            case "a":
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 int8 = -1
                    var v26 uint64 = 2
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            default:
                switch fi10 {
                case 545:
                    for fi27 := 0; fi27 < 6; fi27++ {
                        v28 := v1[fi27]
                        v29 := []int{fi10, fi10, v28, v3}
                        if false {
                            var v30 uint32 = 1
                            switch (v9 + "foo bar") {
                            case "":
                                v31 := GenS2{F0: 1, F1: 982, F2: 9223372036854775807}
                                _ = v31
                            default:
                                fmt.Println(v30, v4, v30)
                            }
                            v4 += (v4 << ((4294967295) & 63))
                            delete(v1, (fi27 << ((772) & 63)))
                            _ = v30
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                case 128:
                    v32 := GenS2{F0: 0, F1: 42, F2: 0}
                    _ = v32
                case 127:
                    v33 := v1[v3]
                    _ = v33
                default:
                    v34 := GenS3{F0: 65535}
                    _ = v34
                }
            }
        }
        if (!true) {
            v35 := GenS4{F0: (-1 != 100), F1: 4294967295}
            _ = v35
        }
        _ = fi10
    }
    var v36 string = v9
    var v37 bool = false
    switch (v3 | -9223372036854775808) {
    case -1:
        var v38 float32 = (1e10 + ((1.5 - 3.14159) / (0.1 / 1e10)))
        _ = v38
    case 128:
        switch v36 {
        case "hello":
            var v39 float32 = -2.25
            _ = v39
        case "foo bar":
            var v40 float32 = -2.25
            _ = v40
        }
    case -359:
        var v41 uint64 = 4294967295
        _ = v41
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        var v43 int8 = 127
        _ = fi42
        _ = v43
    }
    var v44 string = (((v36 + v5) + v36) + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v8
    _ = v9
}
