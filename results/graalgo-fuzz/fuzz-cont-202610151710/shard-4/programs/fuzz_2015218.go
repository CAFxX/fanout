package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int16, p1 int32, p2 uint16) int16 {
    for fi5 := 0; fi5 < 6; fi5++ {
        p2 = 691
        _ = fi5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        p1 -= p1
        v7 := []int{fi6, 2147483647, (fi6 % (((-(fi6))) | 1))}
        _ = fi6
        _ = v7
    }
    v8 := []int{174}
    var v9 uint8 = 64
    var v10 string = ("\n\t" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint64 = 1
    var v4 int32 = 1
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int = fi11
        v13 := GenS1{F0: 1}
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        if true {
            if true {
                var v15 float32 = 0.1
                fmt.Println(v4, v3, v1)
                _ = v15
            } else {
                delete(v2, (^(v1)))
                v16 := []int{}
                var v17 int = (fi14 ^ 0)
                _ = v16
                _ = v17
            }
            var v18 string = "foo bar"
            _ = v18
        } else {
            v1 += 42
            v3 -= (v3 << ((42) & 63))
            v3 = (^(v3))
        }
        _ = fi14
    }
    v4 = 0
    v19 := GenS2{F0: ((!("foo bar" == "\n\t")) || true)}
    var v20 int = -9223372036854775808
    if true {
        for fi21 := 0; fi21 < 3; fi21++ {
            v3 -= (v3 << ((42) & 63))
            v22 := GenS1{F0: 0}
            fmt.Println(v4, fi21)
            _ = fi21
            _ = v22
        }
        if (-1 == 360) {
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := GenS3{F0: v1, F1: (^(v3))}
                for fi25 := 0; fi25 < 19; fi25++ {
                    switch "foo bar" {
                    case "hello":
                        var v26 string = (("foo bar" + ("foo bar" + "\n\t")) + ("" + "foo bar"))
                        _ = v26
                    case "foo bar":
                        v27 := GenS4{F0: 42, F1: 32767}
                        _ = v27
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            if false {
                                v28 := GenS3{F0: v1, F1: (v3 ^ v3)}
                                v1 -= v1
                                v1 = (^(v1))
                                _ = v28
                            }
                            fmt.Println(v3, fi25, v4)
                            if true {
                                _ = len(v2)
                            }
                        case "hello":
                            switch (fi23 | 32768) {
                            case 1:
                                v1 += v1
                            }
                        }
                    }
                    _ = fi25
                }
                _ = fi23
                _ = v24
            }
            var v29 uint16 = 633
            v3 += (v3 ^ v3)
            _ = v29
        } else {
            var v30 float32 = -2.25
            v4 = -1
            _ = v30
        }
        if (!(!(true || (false || false)))) {
            for fi31 := 0; fi31 < 11; fi31++ {
                var v32 string = "a"
                var v33 uint64 = v3
                var v34 uint64 = v33
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            if false {
                v35 := GenS3{F0: (v1 - 322), F1: 4294967295}
                var v36 int16 = 0
                _ = v35
                _ = v36
            } else {
                v1 += (v1 >> ((0) & 63))
            }
        } else {
            switch ("a" + "\n\t") {
            case "a":
                var v37 int8 = 69
                _ = v37
            case "hello":
                var v38 uint = 1
                _ = v38
            case "foo bar":
                v2[v1] = false
            }
        }
    }
    fmt.Println(v1)
    v4 = (v4 ^ v4)
    v3 -= 844
    var v39 int32 = -1
    fmt.Println(v20)
    v40 := GenS2{F0: ((-2.25 * 0.0) < 1.5)}
    v41 := GenS1{F0: genFn0(genFn0(0, v4, 42), int32(v1), 117)}
    v42 := GenS5{F0: v1}
    switch v20 {
    case 32768:
        fmt.Println(v39, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
