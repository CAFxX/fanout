package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
func genFn0() string {
    var v5 int16 = -32768
    if true {
        v6 := []int{}
        _ = v6
    }
    v7 := []int{}
    var v8 int = 42
    if len(v7) > 0 {
        v7[0] = (^(v8))
    }
    _ = v5
    _ = v7
    _ = v8
    return (("a" + "hello") + "foo bar")
}

func main() {
    v1 := []int{127}
    v2 := make(map[string]int)
    v3 := []string{"", "", "foo bar", "hello"}
    v4 := make(map[uint32]uint32)
    var v9 bool = true
    switch 32767 {
    case -1:
        var v10 int64 = -581
        _ = v10
    case 51:
        v9 = true
    default:
        fmt.Println(v9)
    }
    if (!(4294967295 >= 4294967295)) {
        var v11 uint64 = 1
        switch -993 {
        case 42:
            v9 = (-73 >= 0)
        default:
            v9 = v9
        }
        _ = v11
    } else {
        switch -437 {
        case 32767:
            var v12 uint16 = 760
            _ = v12
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v13 := v2[("a" + "hello")]
    v9 = ((("\n\t" + "") >= "hello") && (!true))
    v13 = v13
    for fi14 := 0; fi14 < 12; fi14++ {
        if v9 {
            var v15 uint16 = 2
            _ = v15
        } else {
            v16 := GenS1{F0: 2, F1: 1, F2: 1}
            v13 -= 42
            if v9 {
                v17 := GenS2{F0: 0}
                switch "" {
                case "a":
                    var v18 bool = v9
                    _ = v18
                }
                v3 = append(v3, "")
                _ = v17
            } else {
                var v19 uint = 2
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 float64 = 1.5
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            _ = v16
        }
        var v22 uint8 = 42
        var v23 uint16 = 363
        _ = fi14
        _ = v22
        _ = v23
    }
    v9 = false
    _ = len(v1)
    v24 := GenS3{F0: 0, F1: (!false)}
    v9 = false
    v25 := GenS4{F0: 2, F1: -927}
    for fi26 := 0; fi26 < 6; fi26++ {
        switch (-(v13)) {
        case 128:
            v27 := GenS1{F0: 255, F1: 127, F2: 100}
            _ = v27
        case 467:
            v9 = ((fi26 % (((fi26 & v13)) | 1)) < (^(fi26)))
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            v29 := GenS3{F0: -358, F1: true}
            _ = fi28
            _ = v29
        }
        _ = fi26
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 9; fi31++ {
            switch "" {
            case "a":
                if v9 {
                    v13 = 650
                    v13 -= fi31
                } else {
                    v32 := GenS4{F0: 42, F1: 42}
                    if v9 {
                        v13 = (v13 << ((2) & 63))
                        v13 += fi31
                    }
                    _ = v32
                }
            case "hello":
                v9 = (0 < -883)
            }
            v13 = (fi30 * fi31)
            var v33 bool = true
            _ = fi31
            _ = v33
        }
        v9 = v9
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 uint32 = 4294967295
            if true {
                var v36 string = ((((("hello" + "\n\t") + ("\n\t" + "a")) + ("hello" + ("foo bar" + "\n\t"))) + "\n\t") + "a")
                v35 = v35
                fmt.Println(v36, v35)
                _ = v36
            } else {
                v37 := GenS5{F0: 65535}
                v38 := v2["hello"]
                _ = v37
                _ = v38
            }
            v39 := GenS6{F0: 255}
            v13 += 4294967295
            _ = fi34
            _ = v35
            _ = v39
        }
        _ = fi30
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
