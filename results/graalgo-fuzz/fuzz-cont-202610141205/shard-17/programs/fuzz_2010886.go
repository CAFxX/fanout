package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int }
type GenS2 struct { F0 uint32; F1 bool; F2 int32 }
type GenS3 struct { F0 int32; F1 bool; F2 int }
type GenS4 struct { F0 int16; F1 bool; F2 uint32 }
type GenS5 struct { F0 int }
func genFn0() bool {
    v5 := GenS1{F0: 0, F1: 435, F2: 128}
    var v6 int32 = -1
    switch 1 {
    case 32767:
        v7 := []int{-2147483648, -2147483648}
        _ = v7
    }
    _ = v5
    _ = v6
    return (99 < 225)
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 65535
    var v3 int16 = 1
    var v4 string = ("hello" + "a")
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint16 = 0
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 710
    var v11 int8 = 0
    var v12 string = ((v4 + v4) + v4)
    var v13 float64 = (3.14159 + 0.0)
    v14 := GenS2{F0: 803, F1: (v12 <= v12), F2: -2147483648}
    switch -275 {
    case -9223372036854775808:
        var v15 bool = (32767 != 65535)
        _ = v15
    case -1:
        for fi16 := 0; fi16 < 18; fi16++ {
            switch -201 {
            case 128:
                var v17 bool = (!(int(v10) < (-(fi16))))
                _ = v17
            case 2147483647:
                v18 := GenS1{F0: 106, F1: v10, F2: fi16}
                _ = v18
            case 127:
                v4 = (v4 + (("foo bar" + v4) + ("foo bar" + (v12 + "a"))))
            default:
                v19 := []int{(fi16 ^ (-(fi16)))}
                _ = v19
            }
            fmt.Println(v11)
            var v20 uint32 = (v10 >> (((v2 >> (((^(v2))) & 63))) & 63))
            _ = fi16
            _ = v20
        }
    case 884:
        fmt.Println(v12, v10)
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := []int{(fi21 * fi21), 42, fi21}
        _ = fi21
        _ = v22
    }
    if false {
        v2 = (v2 % ((255) | 1))
        v23 := []int{-9223372036854775808}
        var v24 float64 = v13
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        if true {
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 int32 = 2147483647
                _ = fi26
                _ = v27
            }
        }
        _ = fi25
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        switch v12 {
        case "foo bar":
            v11 = int8(v3)
        case "hello":
            for fi29 := 0; fi29 < 11; fi29++ {
                switch (fi28 + fi28) {
                case 675:
                    v30 := GenS3{F0: int32(fi28), F1: true, F2: int(v10)}
                    v4 = "foo bar"
                    _ = v30
                case 4294967295:
                    fmt.Println(fi29)
                }
                v31 := []int{}
                v32 := GenS4{F0: -32768, F1: false, F2: v10}
                _ = fi29
                _ = v31
                _ = v32
            }
        case "":
            var v33 string = v12
            _ = v33
        }
        _ = fi28
    }
    if false {
        var v34 string = ("hello" + (v12 + ""))
        v35 := []int{786, int(v11), 0, -733}
        _ = v34
        _ = v35
    } else {
        v36 := GenS2{F0: v10, F1: true, F2: 1}
        v37 := GenS5{F0: 32767}
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 uint32 = 993
        fmt.Println(v10)
        for fi40 := 0; fi40 < 8; fi40++ {
            switch ("" + v12) {
            case "a":
                if true {
                    var v41 bool = (v11 == (v11 ^ int8(v39)))
                    var v42 float32 = 0.1
                    _ = v41
                    _ = v42
                } else {
                    var v43 int32 = -1
                    _ = v43
                }
            default:
                switch (v4 + "\n\t") {
                case "hello":
                    v10 -= v39
                default:
                    var v44 uint32 = v10
                    _ = v44
                }
            }
            var v45 uint = v2
            var v46 uint8 = 128
            _ = fi40
            _ = v45
            _ = v46
        }
        _ = fi38
        _ = v39
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        var v48 uint32 = 2147483648
        for fi49 := 0; fi49 < 1; fi49++ {
            if true {
                v50 := GenS2{F0: v10, F1: ((0 != -1) && (!false)), F2: -2147483648}
                switch ("foo bar" + v12) {
                case "":
                    v51 := []int{fi49}
                    _ = v51
                case "a":
                    v52 := v1[(v12 + "hello")]
                    _ = v52
                default:
                    switch (fi47 + fi49) {
                    case 32768:
                        v1["a"] = "foo bar"
                    case -2147483648:
                        var v53 int16 = v3
                        _ = v53
                    }
                }
                v54 := GenS2{F0: 1, F1: true, F2: 0}
                v55 := []int{-1, 9223372036854775807}
                _ = v50
                _ = v54
                _ = v55
            } else {
                v56 := []int{(fi47 << ((v2) & 63)), (-(fi47)), fi47}
                _ = v56
            }
            _ = fi49
        }
        if false {
            var v57 uint32 = 42
            fmt.Println(fi47)
            v2 = v2
            _ = v57
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
