package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint; F1 uint16 }
func genFn0() int64 {
    if (32767 == 683) {
        var v5 string = "hello"
        _ = v5
    } else {
        var v6 uint32 = 596
        _ = v6
    }
    if false {
        v7 := []int{-2147483648, 32767, 43}
        var v8 uint16 = 1
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: "a"}
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{}
        var v12 float64 = 0.0
        if (42 == -2147483648) {
            v12 = ((v12 * 1e10) - v12)
        } else {
            v12 = 3.14159
            var v13 int64 = 382
            for fi14 := 0; fi14 < 13; fi14++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi14
            }
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v9
    return 1
}

func main() {
    var v1 bool = (0 >= 793)
    v2 := make(map[string]bool)
    var v3 uint16 = 155
    var v4 int16 = 1000
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v3)
        _ = fi15
    }
    v16 := GenS1{F0: 103, F1: ("hello" + "hello")}
    if v1 {
        if (!(((float64(((1e10 + 1.5) / 1e10)) * (-2.25 + 1.5)) > 1e10) || (!true))) {
            fmt.Println(v3, v4)
            var v17 int16 = v4
            _ = v17
        } else {
            v18 := []int{-180}
            _ = v18
        }
    }
    fmt.Println(v3, v4)
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int32 = -1
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 42, F1: ""}
    for fi22 := 0; fi22 < 1; fi22++ {
        switch fi22 {
        case -205:
            if (!(-128 < 127)) {
                v1 = v1
            }
        case 1:
            v23 := []int{fi22, 32767}
            _ = v23
        default:
            v24 := []int{fi22}
            _ = v24
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            if false {
                if v1 {
                    v3 -= 2
                    v1 = (!(!true))
                    fmt.Println(fi22)
                } else {
                    v26 := GenS1{F0: -2147483648, F1: "a"}
                    switch "\n\t" {
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            v27 := []int{fi25}
                            _ = v27
                        }
                        var v28 int = -550
                        if (!(v1 && v1)) {
                            delete(v2, "foo bar")
                            v29 := GenS2{F0: 31, F1: v3}
                            var v30 int64 = 682
                            v1 = true
                            _ = v29
                            _ = v30
                        } else {
                            v31 := []int{}
                            _ = v31
                        }
                        _ = v28
                    }
                    _ = v26
                }
                if v1 {
                    v1 = v1
                    v32 := v2[""]
                    v2["foo bar"] = (127 >= 100)
                    _ = v32
                }
            }
            if ((!false) && v1) {
                fmt.Println(v4, v1, v3)
                fmt.Println(v3, v4)
                if (v1 && (v4 == v4)) {
                    var v33 uint32 = 4294967295
                    var v34 int = (fi22 + (^(fi25)))
                    _ = v33
                    _ = v34
                }
            } else {
                if v1 {
                    fmt.Println(fi25, fi22)
                    delete(v2, ("a" + "\n\t"))
                } else {
                    v1 = v1
                    fmt.Println(v3)
                }
                if ((!true) && false) {
                    var v35 uint32 = 4294967295
                    var v36 int64 = 0
                    _ = v35
                    _ = v36
                } else {
                    var v37 int32 = 42
                    v38 := GenS2{F0: 42, F1: (v3 | 65535)}
                    _ = v37
                    _ = v38
                }
            }
            var v39 int = (fi25 << ((0) & 63))
            _ = fi25
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            if (0 >= uint32(v3)) {
                var v41 float64 = (0.0 + (3.14159 + 1e10))
                _ = v41
            }
            _ = fi40
        }
        _ = fi22
    }
    var v42 uint = 1
    if (v1 || (862 > 42)) {
        for fi43 := 0; fi43 < 12; fi43++ {
            v44 := GenS1{F0: 1, F1: "hello"}
            v45 := GenS1{F0: 2147483647, F1: ("hello" + "\n\t")}
            var v46 uint = v42
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        v4 += (v4 / ((42) | 1))
        for fi47 := 0; fi47 < 19; fi47++ {
            var v48 uint = (v42 - (v42 + v42))
            for fi49 := 0; fi49 < 6; fi49++ {
                if (1 < int32(v48)) {
                    var v50 int8 = 0
                    var v51 uint64 = 647
                    _ = v50
                    _ = v51
                }
                v1 = false
                if (1 >= 0) {
                    var v52 int16 = 645
                    _ = v52
                } else {
                    v3 = 369
                    _ = len(v2)
                }
                _ = fi49
            }
            v53 := []int{-496, fi47}
            _ = fi47
            _ = v48
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v42
}
