package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool) int16 {
    var v6 int32 = 42
    p0 = p0
    switch "foo bar" {
    case "hello":
        var v7 string = ("hello" + "hello")
        _ = v7
    case "foo bar":
        var v8 int16 = -32768
        _ = v8
    default:
        if p0 {
            v9 := GenS1{F0: 376, F1: 100}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v6 = (-(v6))
        var v11 uint8 = 163
        _ = fi10
        _ = v11
    }
    if (!(!(!p0))) {
        var v12 int32 = 1
        v6 = (v12 & -1)
        var v13 float64 = 0.0
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1() int8 {
    var v14 int8 = 127
    if true {
        fmt.Println(v14)
    }
    if false {
        v15 := []int{-297}
        _ = v15
    }
    _ = v14
    return (v14 & (v14 / ((81) | 1)))
}

func main() {
    var v1 int32 = -768
    v2 := make(map[int]uint32)
    v3 := []uint32{4294967295, 65535, 2, 2147483648}
    var v4 uint64 = 1
    var v5 float64 = ((-2.25 + -2.25) - 1e10)
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v16 uint8 = 127
        switch "a" {
        case "hello":
            v16 += v16
        }
        _ = v16
    } else {
        fmt.Println(v5, v4, v1)
        switch ("\n\t" + "a") {
        case "a":
            var v17 string = ("a" + ("hello" + ("a" + "hello")))
            _ = v17
        case "":
            v18 := GenS1{F0: 992, F1: 1}
            _ = v18
        case "foo bar":
            v5 = 0.1
        }
    }
    _ = len(v3)
    if false {
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 string = ("a" + "a")
            v21 := GenS2{F0: false}
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            switch fi22 {
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            case 128:
                if len(v3) > 0 {
                    v3[0] = 42
                }
            default:
                var v23 uint32 = 2
                _ = v23
            }
            _ = fi22
        }
        if (v5 >= ((-2.25 * v5) * v5)) {
            fmt.Println(v5, v4)
            v4 = 892
            var v24 int32 = v1
            _ = v24
        } else {
            var v25 bool = false
            fmt.Println(v5)
            fmt.Println(v25, v25, v25)
            _ = v25
        }
        v26 := v2[0]
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS2{F0: ((fi27 + -2147483648) <= fi27)}
        if true {
            if ((!true) && ("hello" < "")) {
                v29 := GenS3{F0: (v4 | v4)}
                switch (-(fi27)) {
                case 0:
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v4 += 42
                        _ = fi30
                    }
                case 2147483647:
                    var v31 int16 = int16(v4)
                    _ = v31
                }
                _ = v29
            } else {
                v32 := GenS2{F0: ("\n\t" >= "a")}
                _ = v32
            }
        } else {
            fmt.Println(v4, v4, v4)
            if false {
                if true {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        var v34 bool = false
                        _ = fi33
                        _ = v34
                    }
                    var v35 uint32 = 2147483648
                    _ = v35
                } else {
                    switch -876 {
                    case 4294967295:
                        var v36 int8 = 1
                        _ = v36
                    case -925:
                        v4 += (^(v4))
                    case -754:
                        fmt.Println(v1, v5, v1)
                    }
                }
            } else {
                var v37 int16 = genFn0(true)
                _ = v37
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                delete(v2, -9223372036854775808)
                var v39 bool = false
                _ = fi38
                _ = v39
            }
        }
        _ = fi27
        _ = v28
    }
    v40 := v2[4294967295]
    v41 := GenS2{F0: (0.0 > 1.5)}
    fmt.Println(v5, v40)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
