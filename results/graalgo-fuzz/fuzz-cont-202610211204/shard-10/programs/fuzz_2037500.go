package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0(p0 int) uint32 {
    if false {
        v6 := GenS1{F0: 1, F1: (true || (!(true || false))), F2: 42}
        switch (p0 / ((-940) | 1)) {
        case 974:
            var v7 uint8 = 128
            _ = v7
        case 42:
            v8 := GenS1{F0: 0, F1: ((!(0 <= 2147483647)) && true), F2: 311}
            _ = v8
        case 127:
            v9 := []int{p0}
            _ = v9
        }
        v10 := []int{p0, (p0 | 647), -533}
        _ = v10
        _ = v6
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        p0 += (p0 >> ((800) & 63))
        switch 310 {
        case 1:
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 uint16 = 492
                var v14 uint8 = 128
                _ = fi12
                _ = v13
                _ = v14
            }
        case 42:
            var v15 int = fi11
            _ = v15
        case -2147483648:
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 float64 = (((3.14159 * float64((1.5 / 1e10))) - 3.14159) * -2.25)
                var v18 int8 = 127
                fmt.Println(fi11, fi16)
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        if false {
            switch "a" {
            case "foo bar":
                if (!false) {
                    v19 := []int{32767, 9223372036854775807}
                    switch (-(p0)) {
                    case 4294967295:
                        v20 := GenS1{F0: 57, F1: false, F2: 0}
                        _ = v20
                    case 1:
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    default:
                        for fi21 := 0; fi21 < 3; fi21++ {
                            p0 = (fi11 & fi11)
                            _ = fi21
                        }
                    }
                    _ = v19
                }
            default:
                v22 := []int{}
                _ = v22
            }
            fmt.Println(fi11, p0)
            var v23 uint64 = 322
            _ = v23
        }
        _ = fi11
    }
    var v24 uint64 = 551
    var v25 string = "hello"
    _ = p0
    _ = v24
    _ = v25
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    var v4 bool = (!true)
    v5 := make(map[int]bool)
    switch "" {
    case "\n\t":
        switch ("hello" + ("a" + "")) {
        case "":
            var v26 uint8 = 255
            _ = v26
        case "hello":
            v27 := GenS1{F0: 255, F1: (-1 > 42), F2: 0}
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 14; fi28++ {
                v4 = (true || ((!(false && (("\n\t" + ("foo bar" + "hello")) <= "a"))) && (v4 && v4)))
                _ = fi28
            }
        default:
            v29 := GenS2{F0: 1, F1: 48, F2: "a"}
            _ = v29
        }
    case "a":
        if (255 > 0) {
            v30 := GenS1{F0: 0, F1: true, F2: 42}
            for fi31 := 0; fi31 < 12; fi31++ {
                if v4 {
                    v32 := []int{}
                    _ = v32
                } else {
                    switch -14 {
                    case 42:
                        var v33 int64 = (v2 << ((0) & 63))
                        _ = v33
                    case -381:
                        switch ("" + "foo bar") {
                        case "\n\t":
                            v4 = (!(v4 && v4))
                        case "foo bar":
                            var v34 uint8 = 241
                            _ = v34
                        }
                    }
                    v4 = (!(908 <= 65535))
                }
                v2 = (v2 & v2)
                _ = fi31
            }
            _ = v30
        }
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        v4 = (!(("foo bar" + "") == ("" + "\n\t")))
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        var v35 string = ("" + "foo bar")
        _ = v35
    }
    var v36 uint16 = 255
    var v37 int32 = 42
    var v38 uint64 = 42
    if v4 {
        v38 = 1000
    } else {
        fmt.Println(v2, v38)
    }
    var v39 string = ""
    if (2 > genFn0(32767)) {
        var v40 float32 = -2.25
        for fi41 := 0; fi41 < 14; fi41++ {
            v42 := GenS3{F0: v36, F1: (-114 <= 100), F2: 32767}
            v39 = ("" + "hello")
            _ = fi41
            _ = v42
        }
        v4 = (!(345 < 390))
        _ = v40
    }
    if v4 {
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 int32 = (v37 * int32(v38))
            _ = fi43
            _ = v44
        }
        var v45 uint16 = 255
        _ = v45
    } else {
        for fi46 := 0; fi46 < 15; fi46++ {
            v47 := GenS3{F0: 255, F1: (!(-2.25 != 1.5)), F2: v36}
            var v48 int8 = 0
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    var v49 int64 = (^(v2))
    if v4 {
        fmt.Println(v49)
        v38 = 18446744073709551615
    }
    var v50 string = ""
    v4 = ((v37 + v37) > (v37 << ((1) & 63)))
    for fi51 := 0; fi51 < 5; fi51++ {
        v52 := GenS3{F0: (v36 / ((v36) | 1)), F1: (!v4), F2: (v36 | v36)}
        _ = fi51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
