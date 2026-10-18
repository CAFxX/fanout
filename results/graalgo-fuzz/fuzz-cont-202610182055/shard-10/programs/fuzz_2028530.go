package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16, p1 int16, p2 int64) int16 {
    switch ("\n\t" + "foo bar") {
    case "":
        v6 := []int{}
        _ = v6
    }
    fmt.Println(p2)
    if false {
        v7 := GenS1{F0: true, F1: (("" + ((((("foo bar" + "\n\t") + "a") + "") + "foo bar") + ("foo bar" + "foo bar"))) == ("" + ("" + "")))}
        if (true || true) {
            v8 := GenS2{F0: -600, F1: 42, F2: 1}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 uint8 = 2
                v10 -= v10
                if false {
                    p0 += (p0 * p0)
                } else {
                    var v11 uint32 = 1
                    v12 := GenS1{F0: (!((v11 == (v11 >> ((634) & 63))) || false)), F1: true}
                    var v13 float32 = (((0.1 - 0.0) * float32(3.14159)) + 1e10)
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            var v14 string = "a"
            p0 = (p0 * (p0 * (p0 % ((p0) | 1))))
            _ = v14
        }
        _ = v7
    }
    var v15 uint = 904
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    var v5 bool = (2147483647 == 1)
    if true {
        v16 := GenS3{F0: 0}
        _ = v16
    } else {
        v17 := GenS4{F0: 1}
        v18 := []int{}
        v5 = ((612 >= 2) && v5)
        _ = v17
        _ = v18
    }
    v3 -= 1e10
    var v19 uint16 = 764
    for fi20 := 0; fi20 < 20; fi20++ {
        fmt.Println(v19, v3)
        fmt.Println(v5, fi20, fi20)
        _ = fi20
    }
    switch 9223372036854775807 {
    case 32767:
        v5 = (!(!(!v5)))
    case 190:
        var v21 uint = 879
        _ = v21
    case 42:
        if (!(0 == 2)) {
            for fi22 := 0; fi22 < 11; fi22++ {
                delete(v2, -2147483648)
                v23 := GenS2{F0: fi22, F1: 923, F2: 0}
                _ = fi22
                _ = v23
            }
            v1[2147483647] = "foo bar"
        } else {
            if (338 <= 750) {
                var v24 uint16 = 32767
                var v25 int64 = int64(v19)
                if (1000 < -22) {
                    var v26 float32 = (3.14159 * v3)
                    var v27 uint = 945
                    for fi28 := 0; fi28 < 6; fi28++ {
                        if v5 {
                            var v29 uint8 = 42
                            _ = v29
                        } else {
                            if ((!true) && false) {
                                v3 = (v26 + ((((((0.0 + float32(0.1)) / (float32((3.14159 / 3.14159)) / 1.0)) + v3) * v3) * v3) - float32(-2.25)))
                            } else {
                                switch ((("hello" + "a") + "hello") + "\n\t") {
                                case "hello":
                                    delete(v4, (fi28 | fi28))
                                case "a":
                                    if (2147483647 <= 42) {
                                        var v30 float64 = -2.25
                                        for fi31 := 0; fi31 < 6; fi31++ {
                                            v26 -= v26
                                            v32 := GenS3{F0: 127}
                                            _ = fi31
                                            _ = v32
                                        }
                                        _ = v30
                                    }
                                case "":
                                    v33 := []int{fi28}
                                    _ = v33
                                }
                                var v34 float32 = v3
                                v35 := GenS2{F0: 9223372036854775807, F1: 0, F2: v25}
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = fi28
                    }
                    _ = v26
                    _ = v27
                }
                _ = v24
                _ = v25
            } else {
                var v36 int16 = 42
                _ = v36
            }
        }
    }
    v37 := v1[2]
    if (v5 || v5) {
        v38 := []int{128, -714, -785, 782}
        switch v37 {
        case "hello":
            v3 += (1e10 + v3)
        case "foo bar":
            v19 = 1
        }
        switch (v37 + ("" + v37)) {
        case "hello":
            _ = len(v38)
        case "":
            if len(v38) > 0 {
                v38[0] = 127
            }
        case "\n\t":
            if true {
                for fi39 := 0; fi39 < 5; fi39++ {
                    v40 := GenS4{F0: -564}
                    v3 = 0.0
                    _ = fi39
                    _ = v40
                }
            }
        }
        _ = v38
    }
    var v41 int16 = 288
    v37 = ("" + "")
    v41 = (v41 + (^(v41)))
    v5 = (2 >= 164)
    var v42 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
