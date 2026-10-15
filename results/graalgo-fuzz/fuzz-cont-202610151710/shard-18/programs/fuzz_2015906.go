package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 bool) bool {
    v5 := []int{9223372036854775807, 0, 32768}
    var v6 string = "hello"
    p0 = true
    _ = p0
    _ = v5
    _ = v6
    return (0 > 0)
}
func genFn1() uint {
    if false {
        var v7 uint32 = 9
        fmt.Println(v7)
        var v8 uint16 = 2
        _ = v7
        _ = v8
    }
    var v9 int64 = -119
    v9 = (-(v9))
    if true {
        v9 = v9
        v9 -= v9
        for fi10 := 0; fi10 < 4; fi10++ {
            if (!(true && false)) {
                v11 := []int{}
                var v12 int16 = -774
                var v13 int = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            switch (^(fi10)) {
            case -982:
                v9 += -392
            }
            _ = fi10
        }
    }
    var v14 uint64 = 849
    _ = v14
    _ = v9
    return 724
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    var v3 int16 = 446
    v4 := make(map[string]int)
    for fi15 := 0; fi15 < 3; fi15++ {
        switch ("a" + "a") {
        case "a":
            for fi16 := 0; fi16 < 20; fi16++ {
                v3 -= v3
                var v17 float64 = -2.25
                _ = fi16
                _ = v17
            }
        case "":
            for fi18 := 0; fi18 < 11; fi18++ {
                switch 2147483647 {
                case 0:
                    v19 := v1[(-(fi18))]
                    _ = v19
                case 1:
                    v1[fi15] = false
                }
                _ = fi18
            }
        default:
            switch "hello" {
            case "a":
                if (-473 < -9223372036854775808) {
                    var v20 int = 938
                    _ = v20
                }
            case "":
                fmt.Println(fi15)
            }
        }
        switch ("foo bar" + "a") {
        case "hello":
            v3 = (v3 & (v3 & (v3 & v3)))
        case "\n\t":
            if (!((0.0 * float32(3.14159)) >= 0.0)) {
                fmt.Println(v3, fi15)
            }
        default:
            var v21 uint16 = 840
            _ = v21
        }
        _ = fi15
    }
    var v22 bool = (!((255 != 42) && (703 >= 756)))
    if (2147483648 <= 2147483647) {
        v23 := GenS1{F0: 42}
        var v24 float64 = ((0.1 - 1.5) * (1e10 / (((3.14159 * 3.14159) - -2.25) + -2.25)))
        _ = v23
        _ = v24
    }
    var v25 uint = 65535
    fmt.Println(v3, v3, v22)
    v26 := v4["a"]
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            v2 = append(v2, "hello")
            if (235 != 42) {
                if v22 {
                    var v29 uint8 = 255
                    switch fi27 {
                    case -673:
                        fmt.Println(v3, v29, v26)
                    default:
                        switch -9223372036854775808 {
                        case 54:
                            switch 9223372036854775807 {
                            case 4294967295:
                                if v22 {
                                    fmt.Println(v26, fi28)
                                } else {
                                    var v30 uint32 = 833
                                    _ = v30
                                }
                            case 247:
                                v31 := GenS1{F0: 1}
                                v32 := GenS1{F0: 128}
                                _ = v31
                                _ = v32
                            case 699:
                                v33 := GenS1{F0: v29}
                                v3 = (-(v3))
                                _ = v33
                            }
                        }
                    }
                    _ = v29
                }
            }
            _ = fi28
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS2{F0: ""}
            _ = fi34
            _ = v35
        }
        _ = fi27
    }
    v3 -= v3
    v36 := GenS2{F0: (("a" + "a") + (("foo bar" + "\n\t") + "hello"))}
    var v37 int64 = -175
    v26 = v26
    v26 -= v26
    if v22 {
        v3 -= v3
        var v38 float64 = (0.0 / 0.1)
        for fi39 := 0; fi39 < 6; fi39++ {
            switch "hello" {
            case "\n\t":
                v40 := GenS3{F0: (v25 << (((v25 & v25)) & 63)), F1: 465}
                _ = v40
            case "":
                v22 = false
            case "hello":
                fmt.Println(v25, v3)
            default:
                var v41 uint = genFn1()
                _ = v41
            }
            var v42 string = ("" + "foo bar")
            for fi43 := 0; fi43 < 11; fi43++ {
                v26 -= -260
                v44 := GenS4{F0: 42}
                _ = fi43
                _ = v44
            }
            _ = fi39
            _ = v42
        }
        _ = v38
    }
    var v45 uint = v25
    v45 = genFn1()
    delete(v1, v26)
    v25 = 0
    if v22 {
        v46 := GenS3{F0: v45, F1: uint32(v45)}
        if false {
            var v47 int32 = -881
            v48 := v4["foo bar"]
            _ = v47
            _ = v48
        } else {
            v3 = v3
            fmt.Println(v45)
            var v49 int16 = v3
            _ = v49
        }
        v50 := GenS1{F0: 18}
        _ = v46
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v45
}
