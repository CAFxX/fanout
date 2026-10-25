package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 int }
type GenS3 struct { F0 uint8 }
func genFn0() string {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 string = "foo bar"
        _ = fi7
        _ = v8
    }
    switch 42 {
    case 4294967295:
        v9 := []int{905, 2147483647, 128}
        _ = v9
    case 32768:
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int32 = -1
            v12 := []int{}
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint8 = 75
                v11 -= (v11 >> ((2) & 63))
                fmt.Println(fi13, v14, v11)
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case -1:
        var v15 uint16 = 2
        _ = v15
    }
    return "\n\t"
}

func main() {
    var v1 int = 32768
    v2 := []int{(^(v1))}
    v3 := []string{"foo bar", ("" + ("hello" + ((("\n\t" + "hello") + "foo bar") + "\n\t"))), ("\n\t" + "foo bar"), "foo bar"}
    v4 := make(map[uint32]string)
    v5 := []string{(((("\n\t" + (((("hello" + "\n\t") + "") + "foo bar") + "a")) + "a") + (("\n\t" + "") + ("\n\t" + ("a" + ("" + ""))))) + ("foo bar" + "\n\t")), "foo bar"}
    var v6 string = ("hello" + (("foo bar" + "") + "foo bar"))
    var v16 int64 = -945
    var v17 int8 = -91
    v17 -= v17
    switch (v1 + 9223372036854775807) {
    case -1:
        if true {
            v18 := GenS1{F0: v1, F1: 4294967295}
            for fi19 := 0; fi19 < 18; fi19++ {
                v2 = append(v2, (v1 * (v1 << ((995) & 63))))
                _ = fi19
            }
            var v20 bool = false
            _ = v18
            _ = v20
        } else {
            if (false && false) {
                if true {
                    _ = len(v2)
                    v2 = append(v2, v1)
                    v6 = (v6 + v6)
                }
                var v21 int64 = v16
                _ = v21
            } else {
                if (v1 > -2147483648) {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v6 = v6
                        _ = fi22
                    }
                    var v23 bool = (0 <= 14)
                    _ = v23
                } else {
                    v1 = (v1 * -1)
                    var v24 string = ("" + ((((v6 + v6) + v6) + v6) + "\n\t"))
                    if (!(v1 != v1)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v25 int64 = (v16 % (((v16 - (v16 | v16))) | 1))
                        v26 := GenS2{F0: v17, F1: 42, F2: (v1 - 1)}
                        _ = v25
                        _ = v26
                    }
                    _ = v24
                }
            }
            for fi27 := 0; fi27 < 18; fi27++ {
                v1 += fi27
                _ = fi27
            }
            if (42 < 521) {
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 string = (v6 + v6)
                    fmt.Println(v16, v17, v17)
                    delete(v4, 0)
                    _ = fi28
                    _ = v29
                }
            } else {
                v17 = (v17 & v17)
                if (42 <= -643) {
                    var v30 uint16 = 1
                    for fi31 := 0; fi31 < 17; fi31++ {
                        var v32 bool = true
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
            }
        }
    }
    v1 += 4294967295
    if (false && (v6 <= "hello")) {
        v33 := GenS3{F0: 128}
        _ = v33
    }
    if (65535 > 2) {
        if ((-2147483648 > 150) && false) {
            for fi34 := 0; fi34 < 3; fi34++ {
                v6 = "foo bar"
                _ = fi34
            }
            var v35 uint64 = 897
            _ = v35
        } else {
            v1 = v1
            v6 = "foo bar"
            v17 = int8(v16)
        }
    }
    v36 := GenS3{F0: 133}
    var v37 uint32 = 0
    for fi38 := 0; fi38 < 6; fi38++ {
        if (2 >= 1) {
            var v39 int16 = 1000
            if true {
                v6 = ("" + v6)
            } else {
                v17 = (v17 + v17)
                if true {
                    var v40 int8 = v17
                    v17 += v40
                    _ = v40
                } else {
                    var v41 int64 = -1
                    var v42 float32 = -2.25
                    v16 = (v41 % (((-(v16))) | 1))
                    _ = v41
                    _ = v42
                }
                var v43 bool = true
                _ = v43
            }
            if true {
                if true {
                    v16 -= (v16 + (v16 / (((v16 / ((9223372036854775807) | 1))) | 1)))
                    v17 += v17
                }
            } else {
                var v44 string = (v6 + v6)
                _ = v44
            }
            _ = v39
        }
        switch (v6 + "a") {
        case "a":
            for fi45 := 0; fi45 < 1; fi45++ {
                var v46 int8 = v17
                var v47 uint = 376
                var v48 int32 = int32(fi38)
                fmt.Println(v48, v46, v1)
                v37 = v37
                var v49 uint32 = 42
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
                _ = v49
            }
        default:
            var v50 int8 = v17
            _ = v50
        }
        _ = fi38
    }
    if (!(750 >= 32767)) {
        v17 -= (-(v17))
        var v51 uint64 = 1
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
